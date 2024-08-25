package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"gorm.io/gorm"
)

func bindArticleApi(app *core.App, r chi.Router) {
	api := articleApi{app}

	r.Get("/", api.list)
	r.Get("/{id}", api.detail)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))

		r.Post("/", api.create)
		r.Put("/{id}", api.update)
		r.Delete("/{id}", api.delete)
	})
}

type articleApi struct {
	*core.App
}

type articleList struct {
	Data      []*models.Article `json:"data"`
	Paginator models.Pagination `json:"paginator"`
}

func (api *articleApi) list(w http.ResponseWriter, r *http.Request) {
	model := api.O.Model(&models.Article{}).Preload("Tags")
	q := r.URL.Query().Get("q")
	var count int64
	println(q)
	page := 1
	pageSize := 10
	if q != "" {
		var params models.SearchURLParams
		err := json.Unmarshal([]byte(q), &params)
		core.P(err)
		page = params.Page + 1
		pageSize = params.PageSize
		// filter
		if params.FilterText != "" {
			model = model.Where(
				"articles.id = ? OR image = ? OR title LIKE ? OR summary LIKE ?",
				append(
					core.CreateSlice[interface{}](2, params.FilterText),
					core.CreateSlice[interface{}](2, fmt.Sprintf("%%%s%%", params.FilterText))...,
				)...,
			)
		}
		for k, v := range params.FilterValues {
			if len(v) == 0 {
				continue
			}
			if k == "tags" {
				model = model.Joins("LEFT JOIN article_tags AS ats ON articles.id = ats.article_id").
					Where("ats.tag_id IN ?", v).
					Group("articles.id").
					Having("COUNT(ats.id) = ?", len(v))
			} else {
				// TODO inject
				model = model.Where(fmt.Sprintf("%s IN ?", k), v)
			}
		}
		err = model.Count(&count).Error
		core.P(err)
		// sort
		if params.SortKey.ID != "" {
			order := "ASC"
			if params.SortKey.Order == "desc" {
				order = "DESC"
			}
			if params.SortKey.ID == "tags" {
				model = model.Select("articles.*, (SELECT COUNT(*) FROM article_tags WHERE article_tags.article_id = articles.id GROUP BY articles.id) AS tag_count").
					Order(fmt.Sprintf("tag_count %s", order))
			} else {
				// TODO 使用反射校验该字段和 Article 结构的json 注释是否匹配 || 找个校验库
				model = model.Order(fmt.Sprintf("%s %s", params.SortKey.ID, order))
			}
		}
	} else {
		err := model.Count(&count).Error
		core.P(err)
	}
	model = model.Order("id DESC")
	var articles []*models.Article
	model = model.Offset((page - 1) * 10).Limit(pageSize)
	err := model.Find(&articles).Error
	core.P(err)
	json.NewEncoder(w).Encode(articleList{
		Data: articles,
		Paginator: models.Pagination{
			Count: int(count),
		},
	})
}

func (api *articleApi) detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.P(err)
	article := new(models.Article)
	err = api.O.Model(&models.Article{}).
		Preload("ArticleContent", func(db *gorm.DB) *gorm.DB {
			return db.Omit("html")
		}).
		Preload("Tags").
		First(article, id).
		Error
	core.P(err)
	json.NewEncoder(w).Encode(article)
}

func (api *articleApi) create(w http.ResponseWriter, r *http.Request) {
	art := new(models.Article)
	err := json.NewDecoder(r.Body).Decode(art)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(art).Error; err != nil {
			return err
		}
		at := &models.ArticleContent{
			ArticleID: art.ID,
		}
		return tx.Create(at).Error
	})
	core.P(err)
	api.JSON(w, art)
}

func (api *articleApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		if v, ok := data["text"]; ok {
			content := &models.ArticleContent{
				ArticleID: uint(id),
				Text:      v.(string),
				// text 字段要和编译好的 html 字段一起
				HTML: data["html"].(string),
			}
			err = tx.Save(content).Error
			if err != nil {
				return err
			}
			data["rune"] = utf8.RuneCountInString(content.Text)
			delete(data, "text")
			delete(data, "html")
		}
		if v, ok := data["tags"]; ok {
			tags := v.([]interface{})
			tx.Where("article_id = ?", id).Delete(new(models.ArticleTags))
			for _, t := range tags {
				artTags := &models.ArticleTags{
					ArticleID: uint(id),
					TagID:     uint(t.(map[string]interface{})["id"].(float64)),
				}
				if err = tx.Create(artTags).Error; err != nil {
					return err
				}
			}
			delete(data, "tags")
		}
		return tx.Model(new(models.Article)).Where("id = ?", id).Updates(data).Error
	})
	core.P(err)
	api.JSON(w, id)
}

func (api *articleApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(new(models.Article), id).Error; err != nil {
			return err
		}
		return tx.Where("article_id = ?", id).Delete(new(models.ArticleTags)).Error
	})
	core.P(err)
	api.JSON(w, id)
}
