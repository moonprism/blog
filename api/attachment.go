package api

import (
	"encoding/json"
	"net/http"

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
