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

func bindGistApi(app *core.App, r chi.Router) {
	api := gistApi{app}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))
		r.Get("/", api.list)
		r.Post("/", api.create)
		r.Put("/{id}", api.update)
		r.Delete("/{id}", api.delete)
	})
}

type gistApi struct {
	*core.App
}

type gistList struct {
	Data       []*models.Gist    `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

func (api *gistApi) list(w http.ResponseWriter, r *http.Request) {
	model := api.O.Model(&models.Gist{})
	page := 1
	pageSize := 10
	var count int64
	q := r.URL.Query().Get("q")
	if q != "" {
		var params models.SearchURLParams[string]
		err := json.Unmarshal([]byte(q), &params)
		core.P(err)
		page = params.Page + 1
		pageSize = params.PageSize
		// filter
		if params.FilterText != "" {
			model = model.Where(
				"`id` = ? OR `title` LIKE ? OR `lang` LIKE ? OR `content` LIKE ?",
				append(
					[]interface{}{params.FilterText},
					core.CreateSlice[interface{}](3, fmt.Sprintf("%%%s%%", params.FilterText))...,
				)...,
			)
		}
		for k, v := range params.FilterValues {
			if len(v) == 0 {
				continue
			}
			if k == "lang" {
				model = model.Where(fmt.Sprintf("%s IN ?", k), v)
			}
		}
		// sort
		if params.SortKey.ID != "" {
			order := "ASC"
			if params.SortKey.Order == "desc" {
				order = "DESC"
			}
			// inject
			model = model.Order(fmt.Sprintf("%s %s", params.SortKey.ID, order))
		}
	}
	err := model.Count(&count).Error
	core.P(err)
	var gists []*models.Gist
	err = model.Order("id DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&gists).
		Error
	core.P(err)
	json.NewEncoder(w).Encode(gistList{
		Data: gists,
		Pagination: models.Pagination{
			Count: int(count),
		},
	})
}

func (api *gistApi) create(w http.ResponseWriter, r *http.Request) {
	gist := new(models.Gist)
	err := json.NewDecoder(r.Body).Decode(gist)
	core.P(err)
	err = api.O.Create(gist).Error
	core.P(err)
	api.JSON(w, gist)
}

func (api *gistApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.P(err)
	gist := new(models.Gist)
	gist.ID = uint(id)
	err = api.O.Model(gist).Updates(data).Error
	core.P(err)
	api.JSON(w, id)
}

func (api *gistApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	gist := new(models.Gist)
	gist.ID = uint(id)
	api.O.Delete(&gist)
	core.P(err)
	api.JSON(w, id)
}
