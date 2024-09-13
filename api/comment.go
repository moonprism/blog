package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
)

func bindCommentApi(app *core.App, r chi.Router) {
	api := commentApi{app}

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
