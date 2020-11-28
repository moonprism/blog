package api

import (
	"git.kicoe.com/blog/write/config"
	"git.kicoe.com/blog/write/errors"
	_ "git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type authResponse struct {
	Token string `json:"token"`
}
// @Summary admin login (jwt)
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginData body service.AuthBody true " "
// @Success 200 {object} authResponse
// @Failure 401 {object} errors.HTTPError "Unauthorized"
// @Router /auth/login [post]
func Login(c echo.Context) error {
	authBody := new(service.AuthBody)
	if err := c.Bind(authBody); err != nil {
		return err
	}

	if service.CheckAuth(authBody) == false {
		return errors.HTTPErrorUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	t, e := token.SignedString([]byte(config.Auth.SigningKey))
	if e != nil {
		return e
	}

	return c.JSON(http.StatusOK, &authResponse{t})
}

// todo
type infoResponse struct {
	IsRegister	bool	`json:"is_register"`
	BackgroundImage	string	`json:"background_image"`
}
// @Summary deprecated
// @bolg system info
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} infoResponse
// @Router /auth/info [get]
func Info(c echo.Context) error {
	return nil
}

// @Summary deprecated
// @admin regist (jwt)
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginData body service.AuthBody true " "
// @Success 200 {object} authResponse
// @Router /auth/regist [post]
func Regist(c echo.Context) error {
	return nil
}
