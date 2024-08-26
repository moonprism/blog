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

func bindAttachmentApi(app *core.App, r chi.Router) {
	api := attachmentApi{app}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))
		r.Get("/", api.list)
		r.Post("/", api.create)
		r.Put("/{id}", api.update)
		r.Delete("/{id}", api.delete)
	})
}

type attachmentApi struct {
	*core.App
}

type Attachment struct {
	models.Attachment
	Year int `json:"year"`
}

type attachmentList struct {
	Data       []*Attachment     `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

func (api *attachmentApi) list(w http.ResponseWriter, r *http.Request) {
	model := api.O.Model(&models.Attachment{})
	page := 1
	pageSize := 20
	var count int64
	q := r.URL.Query().Get("q")
	if q != "" {
		var params models.SearchURLParams
		err := json.Unmarshal([]byte(q), &params)
		core.P(err)
		page = params.Page + 1
		pageSize = params.PageSize
		// filter
		if params.FilterText != "" {
			model = model.Where(
				"`key` = ? OR `summary` LIKE ?",
				params.FilterText,
				fmt.Sprintf("%%%s%%", params.FilterText),
			)
		}
		for k, v := range params.FilterValues {
			if len(v) == 0 {
				continue
			}
			if k == "year" {
				model = model.Where("FROM_UNIXTIME(created, '%Y') IN ?", v)
			}
		}
	}
	err := model.Count(&count).Error
	core.P(err)
	var attachments []*Attachment
	err = model.Select("attachments.*, FROM_UNIXTIME(created, '%Y') AS year").Order("id DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&attachments).
		Error
	core.P(err)
	json.NewEncoder(w).Encode(attachmentList{
		Data: attachments,
		Pagination: models.Pagination{
			Count: int(count),
		},
	})
}

func (api *attachmentApi) create(w http.ResponseWriter, r *http.Request) {
	attachment := new(models.Attachment)
	err := json.NewDecoder(r.Body).Decode(attachment)
	core.P(err)
	err = api.O.Create(attachment).Error
	core.P(err)
	api.JSON(w, attachment)
}

func (api *attachmentApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.P(err)
	attachment := new(models.Attachment)
	attachment.ID = uint(id)
	err = api.O.Model(attachment).Updates(data).Error
	core.P(err)
	api.JSON(w, id)
}

func (api *attachmentApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	attachment := new(models.Attachment)
	attachment.ID = uint(id)
	api.O.Delete(&attachment)
	core.P(err)
	api.JSON(w, id)
}
