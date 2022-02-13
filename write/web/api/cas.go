package api

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"sync/atomic"

	"git.kicoe.com/blog/write/modules/utils"
	"github.com/labstack/echo/v4"
)

type CasAuthorization struct {
	Key string `json:"cas_key"`
}

var casKey atomic.Value

// @Summary get cas key
// @Tags Cas
// @Accept  json
// @Produce  json
// @param Authorization header string true "Authorization"
// @Success 200 {object} CasAuthorization
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Router /cas/key [get]
func CasKey(c echo.Context) error {
	casResponse := new(CasAuthorization)
	key := utils.RandBytes(10)
	casKey.Store(key)
	casResponse.Key = fmt.Sprintf("%x", md5.Sum(key))

	return c.JSON(http.StatusOK, casResponse)
}

// @Summary check cas key
// @Tags Cas
// @Accept  json
// @Produce  json
// @Param key query string true "cas key"
// @Success 200 {string} string "success"
// @Failure 401 {object} errors.HTTPError "Unauthorized"
// @Router /cas/auth [get]
func CasAuth(c echo.Context) error {
	casValue := casKey.Load()
	if casValue == nil || c.QueryParam("key") != fmt.Sprintf("%x", md5.Sum(casValue.([]byte))) {
		return c.String(http.StatusUnauthorized, "401")
	}
	return c.String(http.StatusOK, "success")
}
