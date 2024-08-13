package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/model"
)

func bindAttachmentApi(app *core.App, r chi.Router) {
	api := attachmentApi{app}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))
		r.Get("/", api.list)
		r.Post("/", api.create)
		r.Put("/{id}", api.update)
	})
}

type attachmentApi struct {
	*core.App
}

func (api *attachmentApi) list(w http.ResponseWriter, r *http.Request) {
	var attachments []*model.Attachment
	err := api.O.Model(&model.Attachment{}).Order("id desc").Find(&attachments).Error
	core.P(err)
	json.NewEncoder(w).Encode(&attachments)
}

func (api *attachmentApi) create(w http.ResponseWriter, r *http.Request) {
	attachment := new(model.Attachment)
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
	attachment := new(model.Attachment)
	attachment.ID = uint(id)
	err = api.O.Model(attachment).Updates(data).Error
	core.P(err)
	api.JSON(w, attachment)
}
