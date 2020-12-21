package api

import (
	"github.com/labstack/echo/v4"
	"git.kicoe.com/blog/write/service"
	"git.kicoe.com/blog/write/model"
	"net/http"
)

// @Summary get setting
// @Tags Setting
// @Produce  json
// @param Authorization header string true "Authorization"
// @Success 200 {object} model.Setting
// @Failure 403 {object} errors.HTTPError
// @Router /setting [get]
func SettingDetail(c echo.Context) error {
	setting, _ := service.GetSetting()
	return c.JSON(http.StatusOK, setting)
}

// @Summary set setting
// @Tags Setting
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param settingInfo body model.Setting true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /setting [put]
func UpdateSetting(c echo.Context) error {
	setting := new(model.Setting)
	if err := c.Bind(setting); err != nil {
		return err
	}
	service.SetSetting(setting)
	return c.JSON(http.StatusOK, "ok")
}
