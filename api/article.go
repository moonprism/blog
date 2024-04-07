package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/model"
)

func bindArticleApi(app *core.App, r chi.Router) {
	api := articleApi{app}

	r.Get("/", api.list)
	r.Post("/", api.create)
}

type articleApi struct {
	*core.App
}

func (api *articleApi) list(w http.ResponseWriter, r *http.Request) {
	articles := make([]*model.Article, 0)
	api.O.Find(&articles)
	spew.Dump(articles)
	json.NewEncoder(w).Encode(articles)
}

func (api *articleApi) create(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1)
	status, err := strconv.Atoi(r.PostFormValue("status"))
	if err != nil {
		panic(err)
	}
	article := model.Article{
		Title:   r.PostFormValue("title"),
		Status:  status,
		Image:   "",
		Summary: "sbsb",
		Content: "#I'm Sbsb",
	}

	_, err = api.O.Insert(&article)
	if err != nil {
		spew.Dump(err)
		// w.Write([]byte("xxx"))
	}

}
