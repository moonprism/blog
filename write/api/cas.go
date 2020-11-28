package api

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"git.kicoe.com/blog/write/config"
	"github.com/labstack/echo/v4"
)

type casAuthorization struct {
	Key string `json:"cas_key"`
}

// @Summary get cas key
// @Tags Cas
// @Accept  json
// @Produce  json
// @param Authorization header string true "Authorization"
// @Success 200 {object} casAuthorization
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Router /cas/key [get]
func CasKey(c echo.Context) error {
	casResponse := new(casAuthorization)
	casResponse.Key = fmt.Sprintf("%x", md5.Sum([]byte(config.CAS.Key)))

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
	if c.QueryParam("key") == fmt.Sprintf("%x", md5.Sum([]byte(config.CAS.Key))) {
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusUnauthorized, "401")
}
