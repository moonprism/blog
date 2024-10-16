package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"gorm.io/gorm"
)

func bindTagApi(app *core.App, r chi.Router) {
	api := tagApi{app}

	r.Get("/", api.list)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))

		r.Post("/", api.create)
		r.Put("/{id}", api.update)
		r.Delete("/{id}", api.delete)
	})
}

type tagApi struct {
	*core.App
}

type tagList struct {
	Data []*models.Tag `json:"data"`
}

func (api *tagApi) list(w http.ResponseWriter, r *http.Request) {
	tags := make([]*models.Tag, 0)
	err := api.O.Order("id desc").Find(&tags).Error
	core.P(err)
	api.JSON(w, tagList{
		Data: tags,
	})
}

func (api *tagApi) create(w http.ResponseWriter, r *http.Request) {
	tag := new(models.Tag)
	err := json.NewDecoder(r.Body).Decode(tag)
	core.P(err)
	err = api.O.Create(tag).Error
	core.P(err)
	api.JSON(w, tag)
}

func (api *tagApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.P(err)
	tag := new(models.Tag)
	tag.ID = uint(id)
	err = api.O.Model(tag).Updates(data).Error
	core.P(err)
	api.JSON(w, id)
}

func (api *tagApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(new(models.Tag), id).Error; err != nil {
			return err
		}
		return tx.Where("tag_id = ?", id).Delete(new(models.ArticleTags)).Error
	})
	core.P(err)
	api.JSON(w, id)
}
