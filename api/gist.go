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
	"gorm.io/gorm"
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
	model := api.O.Model(&models.Gist{}).Preload("GistOutput")
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
	_, err = api.O.SqliteFtsDB.Exec(
		"INSERT INTO gists_fts (rowid, title, lang, content) VALUES (?, ?, ?, ?)",
		gist.ID,
		gist.Title,
		gist.Lang,
		gist.Content,
	)
	core.P(err)
	api.JSON(w, gist)
}

func (api *gistApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		gist := new(models.Gist)
		gist.ID = uint(id)

		if v, ok := data["html"]; ok {
			content := &models.GistOutput{
				GistID: gist.ID,
				HTML:   v.(string),
			}
			err = tx.Save(content).Error
			if err != nil {
				return err
			}
			delete(data, "html")
		}

		if len(data) == 0 {
			return nil
		}
		return tx.Model(gist).Updates(data).Error
	})
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

func gistsPageRoute(app *core.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := app.HTML(w, "gists", getAppSettings(app))
		core.P(err)
	}
}

func gistsSearchRoute(app *core.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		rows, err := app.O.SqliteFtsDB.Query(
			`select
				rowid,
				simple_highlight(gists_fts, 0, '[', ']'),
				simple_highlight(gists_fts, 1, '[', ']'),
				simple_highlight(gists_fts, 2, '[', ']')
			from gists_fts where content match simple_query(?) or title match simple_query(?) or lang match ?`,
			core.CreateSlice[interface{}](3, q)...,
		)
		core.P(err)
		defer rows.Close()
		for rows.Next() {
			var id int
			var title, lang, content string
			rows.Scan(&id, &title, &lang, &content)
			fmt.Printf("%d: \n %s \n %s \n %s \n --- \n", id, title, lang, content)
		}
		//app.JSON(w, rows)
	}
}
