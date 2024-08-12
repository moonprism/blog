package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/model"
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
	var article []*model.Article
	err := api.O.Model(&model.Article{}).Preload("Tags").Order("id desc").Find(&article).Error
	core.P(err)
	json.NewEncoder(w).Encode(&article)
}

func (api *articleApi) detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.P(err)
	article := new(model.Article)
	err = api.O.Model(&model.Article{}).
		Preload("ArticleText").
		Preload("Tags").
		First(article, id).
		Error
	core.P(err)
	json.NewEncoder(w).Encode(article)
}

func (api *articleApi) create(w http.ResponseWriter, r *http.Request) {
	art := new(model.Article)
	err := json.NewDecoder(r.Body).Decode(art)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(art).Error; err != nil {
			return err
		}
		// gorm auto!!!
		// for _, t := range art.Tags {
		// 	artTags := &model.ArticleTags{
		// 		ArticleID: art.ID,
		// 		TagID:     t.ID,
		// 	}
		// 	if err = tx.Create(artTags).Error; err != nil {
		// 		return err
		// 	}
		// }
		at := &model.ArticleText{
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
		if v, ok := data["content"]; ok {
			artText := &model.ArticleText{
				ArticleID: uint(id),
				Content:   v.(string),
			}
			err = tx.Save(artText).Error
			if err != nil {
				return err
			}
			data["rune"] = utf8.RuneCountInString(artText.Content)
			delete(data, "content")
		}
		if v, ok := data["tags"]; ok {
			tags := v.([]interface{})
			tx.Where("article_id = ?", id).Delete(new(model.ArticleTags))
			for _, t := range tags {
				artTags := &model.ArticleTags{
					ArticleID: uint(id),
					TagID:     uint(t.(map[string]interface{})["id"].(float64)),
				}
				if err = tx.Create(artTags).Error; err != nil {
					return err
				}
			}
			delete(data, "tags")
		}
		return tx.Model(new(model.Article)).Where("id = ?", id).Updates(data).Error
	})
	core.P(err)
	api.JSON(w, id)
}

func (api *articleApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(new(model.Article), id).Error; err != nil {
			return err
		}
		return tx.Where("article_id = ?", id).Delete(new(model.ArticleTags)).Error
	})
	core.P(err)
}
