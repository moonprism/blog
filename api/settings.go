package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"gorm.io/gorm"
)

func bindSettingsApi(app *core.App, r chi.Router) {
	api := settingApi{app}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))
		r.Get("/", api.detail)
		r.Post("/", api.update)
	})
}

type settingApi struct {
	*core.App
}

type AppSettings struct {
	*core.SystemSet
	*models.Settings
	Pages *[4]string
}

var appSettings = AppSettings{
	Pages: &[4]string{"posts", "gists", "links", "about"},
}

// 简单缓存
func getAppSettings(app *core.App) *AppSettings {
	if appSettings.SystemSet == nil {
		var settings models.Settings
		app.O.First(&settings)
		appSettings.SystemSet = &app.Settings.System
		appSettings.Settings = &settings
	}
	return &appSettings
}

type settingsResponse struct {
	*AppSettings
	LastLoginTime *string `json:"lastLoginTime"`
}

func (api *settingApi) detail(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	username := claims["user"].(string)
	var user models.User
	err := api.O.Where(&models.User{Name: username}).First(&user).Error
	core.P(err)
	lastLoginTime := time.Unix(int64(user.LastLogin), 0).Format("2006年01月02日 15:04")
	api.JSON(w, &settingsResponse{
		getAppSettings(api.App),
		&lastLoginTime,
	})
}

func (api *settingApi) update(w http.ResponseWriter, r *http.Request) {
	var body models.Settings
	err := json.NewDecoder(r.Body).Decode(&body)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		var settings models.Settings
		if err = tx.First(&settings).Error; err != nil {
			if !api.O.IsRecordNotFoundErr(err) {
				return err
			}
			return tx.Create(&settings).Error
		}
		return tx.Model(&models.Settings{}).Where("id = ?", settings.ID).Updates(map[string]interface{}{
			"title":         body.Title,
			"background":    body.Background,
			"margin_bottom": body.MarginBottom,
		}).Error
	})
	core.P(err)
	// no lock
	appSettings.SystemSet = nil
	api.JSON(w, "ok")
}
