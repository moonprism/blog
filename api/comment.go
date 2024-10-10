package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"gorm.io/gorm"
)

func bindCommentApi(app *core.App, r chi.Router) {
	api := commentApi{app}

	r.Post("/", api.post)
	r.Get("/page/{article_id}", api.listOnPage)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))

		r.Get("/", api.list)
		r.Delete("/{id}", api.delete)
	})
}

type commentApi struct {
	*core.App
}

type commentPostError struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Content string `json:"content"`
	System  string `json:"system"`
}

func (api *commentApi) post(w http.ResponseWriter, r *http.Request) {
	comment := new(models.Comment)
	err := json.NewDecoder(r.Body).Decode(comment)
	core.P(err)
	postError := commentPostError{}

	// 或许该找个字段校验库
	if comment.Email == "" || comment.Name == "" || comment.Content == "" {
		return
	}

	if len(comment.Email) > 100 {
		postError.Email = "邮箱过长"
	}
	if len(comment.Name) > 50 {
		postError.Name = "姓名过长"
	}
	if len(comment.Content) > 1000 {
		postError.Content = "文本过长"
	}

	if (postError != commentPostError{}) {
		w.WriteHeader(400)
		api.JSON(w, postError)
		return
	}

	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(comment.Email) {
		w.WriteHeader(400)
		postError.Email = "无效的电子邮箱地址"
		api.JSON(w, postError)
		return
	}

	var count int64
	err = api.O.Model(&models.Article{}).
		Where("id = ?", comment.ArticleID).
		Count(&count).
		Error
	core.P(err)
	if count < 0 {
		w.WriteHeader(422)
		return
	}

	var rootComment models.Comment

	var replyComment models.Comment
	if comment.ReplyCommentID != 0 {
		err = api.O.First(&replyComment, comment.ReplyCommentID).Error
		core.P(err)
		var rootCommentId = replyComment.RootCommentID
		if rootCommentId == 0 {
			rootComment = replyComment
		} else {
			err = api.O.First(&rootComment, rootCommentId).Error
			core.P(err)
		}
		comment.RootCommentID = rootComment.ID
	}

	err = api.O.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(comment).Error; err != nil {
			return err
		}
		if comment.ReplyCommentID != 0 {
			// 更新相关评论时间排序
			rootComment.Updated = comment.Updated
			if err = tx.Save(&rootComment).Error; err != nil {
				return err
			}
		}
		return nil
	})
	core.P(err)
	comment.Email = fmt.Sprintf("%x", md5.Sum([]byte(comment.Email)))
	api.JSON(w, comment)
}

type commentPageItem struct {
	*models.Comment
	SubComments []*models.Comment `json:"sub_comments"`
	HasNext     bool              `json:"has_next"`
}
type commentPageList struct {
	Data    []*commentPageItem `json:"data"`
	HasNext bool               `json:"has_next"`
}

func (api *commentApi) listOnPage(w http.ResponseWriter, r *http.Request) {
	article_id, err := strconv.ParseUint(chi.URLParam(r, "article_id"), 10, 32)
	core.P(err)

	var rootId uint64 = 0
	rootIdParam := r.URL.Query().Get("root_id")
	if rootIdParam != "" {
		rootId, err = strconv.ParseUint(rootIdParam, 10, 32)
		core.P(err)
	}

	var page = 1
	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		core.P(err)
	}

	pageSize := 5
	var count int64
	model := api.O.Model(&models.Comment{}).
		Where("article_id = ?", article_id).
		Where("root_comment_id = ?", rootId)
	err = model.Count(&count).Error
	core.P(err)
	if count == 0 {
		return
	}
	var comments []*models.Comment
	err = model.Order("updated DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&comments).
		Error
	core.P(err)

	pageComments := make([]*commentPageItem, len(comments))
	for i, c := range comments {
		c.Email = fmt.Sprintf("%x", md5.Sum([]byte(c.Email)))
		// 子评论
		var subComments []*models.Comment
		var subCount int64
		if rootId == 0 {
			subModel := api.O.Model(&models.Comment{}).
				Where("root_comment_id = ?", c.ID)
			err = subModel.Count(&subCount).Error
			core.P(err)
			if subCount == 0 {
				pageComments[i] = &commentPageItem{
					Comment:     c,
					SubComments: []*models.Comment{},
					HasNext:     false,
				}
				continue
			}
			err = subModel.Order("id ASC").
				Offset(0).
				Limit(pageSize).
				Find(&subComments).
				Error
			core.P(err)
			for _, c2 := range subComments {
				c2.Email = fmt.Sprintf("%x", md5.Sum([]byte(c2.Email)))
			}
		}
		pageComments[i] = &commentPageItem{
			Comment:     c,
			SubComments: subComments,
			HasNext:     subCount > int64(pageSize*page),
		}
	}
	api.JSON(w, commentPageList{
		Data:    pageComments,
		HasNext: count > int64(pageSize*page),
	})
}

type commentList struct {
	Data       []*models.Comment `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

func (api *commentApi) list(w http.ResponseWriter, r *http.Request) {
	model := api.O.Model(&models.Comment{})
	var count int64
	page := 1
	pageSize := 10
	q := r.URL.Query().Get("q")
	if q != "" {
		var params models.SearchURLParams[int]
		err := json.Unmarshal([]byte(q), &params)
		core.P(err)
		page = params.Page + 1
		pageSize = params.PageSize
		// filter
		if params.FilterText != "" {
			model = model.Where(
				"id = ? OR article_id = ? OR reply_comment_id = ? OR root_comment_id = ? OR `name` LIKE ? OR `email` LIKE ? OR `link` LIKE ? OR `content` LIKE ?",
				append(
					core.CreateSlice[interface{}](4, params.FilterText),
					core.CreateSlice[interface{}](4, fmt.Sprintf("%%%s%%", params.FilterText))...,
				)...,
			)
		}
		err = model.Count(&count).Error
		core.P(err)
		// sort
		if params.SortKey.ID != "" {
			order := "ASC"
			if params.SortKey.Order == "desc" {
				order = "DESC"
			}
			// inject
			model = model.Order(fmt.Sprintf("%s %s", params.SortKey.ID, order))
		}
	} else {
		err := model.Count(&count).Error
		core.P(err)
	}
	var comments []*models.Comment
	err := model.Order("id DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&comments).
		Error
	core.P(err)
	json.NewEncoder(w).Encode(commentList{
		Data: comments,
		Pagination: models.Pagination{
			Count: int(count),
		},
	})
}

func (api *commentApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	comment := new(models.Comment)
	comment.ID = uint(id)
	api.O.Delete(&comment)
	core.P(err)
	api.JSON(w, id)
}
