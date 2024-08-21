package api

import (
	"encoding/json"
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

func (api *articleApi) list(w http.ResponseWriter, r *http.Request) {
	var article []*models.Article
	err := api.O.Model(&models.Article{}).Preload("Tags").Order("id desc").Find(&article).Error
	core.P(err)
	json.NewEncoder(w).Encode(&article)
}

func (api *articleApi) detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.P(err)
	article := new(models.Article)
	err = api.O.Model(&models.Article{}).
		Preload("ArticleContent").
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
