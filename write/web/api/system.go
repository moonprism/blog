package api

import (
	"net/http"

	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/services/comment"
	"git.kicoe.com/blog/write/services/system"
	"github.com/labstack/echo/v4"
)

type SystemInfo struct {
	Theme   models.Theme `json:"theme"`
	Captcha bool         `json:"captcha"`
	Notice  bool         `json:"notice"`
}

// @Summary get setting
// @Tags Setting
// @Produce  json
// @Success 200 {object} SystemInfo
// @Failure 403 {object} errors.HTTPError
// @Router /system/info [get]
func GetSystemInfo(c echo.Context) error {
	theme, _ := system.GetTheme()
	return c.JSON(http.StatusOK, SystemInfo{
		Theme:   theme,
		Captcha: !system.LL.State(),
		Notice:  comment.GetNoticeFlag(),
	})
}

// @Summary set setting
// @Tags Setting
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param settingInfo body models.Theme true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /system/info/theme [put]
func UpdateSystemTheme(c echo.Context) error {
	theme := new(models.Theme)
	if err := c.Bind(theme); err != nil {
		return err
	}
	system.SetTheme(theme)
	return c.JSON(http.StatusOK, "ok")
}

// @Summary set setting
// @Tags Setting
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param settingInfo body SystemInfo true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /system/info/notice [put]
func UpdateSystemNotice(c echo.Context) error {
	info := new(SystemInfo)
	if err := c.Bind(info); err != nil {
		return err
	}
	comment.SetNoticeFlag(info.Notice)
	return c.JSON(http.StatusOK, "ok")
}
