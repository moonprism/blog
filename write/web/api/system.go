package api

import (
	"net/http"

	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/services/system"
	"github.com/labstack/echo/v4"
)

type SystemInfoResponse struct {
	Theme models.Theme `json:"theme"`
	Captcha bool `json:"captcha"`
}


// @Summary get setting
// @Tags Setting
// @Produce  json
// @Success 200 {object} SystemInfoResponse
// @Failure 403 {object} errors.HTTPError
// @Router /system/info [get]
func GetSystemInfo(c echo.Context) error {
	theme, _ := system.GetTheme()
	return c.JSON(http.StatusOK, SystemInfoResponse{
		Theme: theme,
		Captcha: !system.LL.State(),
	})
}

// @Summary set setting
// @Tags Setting
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param settingInfo body models.Theme true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /system/info [put]
func UpdateSystemInfo(c echo.Context) error {
	theme := new(models.Theme)
	if err := c.Bind(theme); err != nil {
		return err
	}
	system.SetTheme(theme)
	return c.JSON(http.StatusOK, "ok")
}
