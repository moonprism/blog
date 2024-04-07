package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/core/http/middleware"
	"github.com/moonprism/blog/model"
)

func bindArticleApi(app *core.App, r chi.Router) {
	api := articleApi{app}

	r.Use(middleware.JsonResponse)
	r.Get("/", api.list)
	r.Get("/{id}", api.detail)
	r.Post("/", api.create)
	r.Put("/{id}", api.update)
	r.Delete("/{id}", api.delete)
}

type articleApi struct {
	*core.App
}

func (api *articleApi) list(w http.ResponseWriter, r *http.Request) {
	articles := make([]*model.Article, 0)
	api.O.Omit("content").Find(&articles)
	json.NewEncoder(w).Encode(&articles)
}

func (api *articleApi) detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.Pf(err)
	article := new(model.Article)
	has, err := api.O.ID(id).Get(article)
	core.Pf(err)
	if !has {
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode(&article)
}

func (api *articleApi) create(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	core.Pf(err)
	article.Rune = utf8.RuneCountInString(article.Content)
	_, err = api.O.Insert(&article)
	core.Pf(err)
}

func (api *articleApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.Pf(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.Pf(err)
	if v, ok := data["content"]; ok {
		data["rune"] = utf8.RuneCountInString(v.(string))
	}
	_, err = api.O.Table(new(model.Article)).ID(id).Update(data)
	core.Pf(err)
}

func (api *articleApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.Pf(err)
	_, err = api.O.ID(id).Delete(new(model.Article))
	core.Pf(err)
}
