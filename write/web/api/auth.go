package api

import (
	"net/http"
	"time"

	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/services/auth"
	"git.kicoe.com/blog/write/services/system"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AuthRequest struct {
	auth.AuthBody
	Captcha string `json:"captcha"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

// @Summary admin login (jwt)
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginData body AuthRequest true " "
// @Success 200 {object} AuthResponse
// @Failure 401 {object} errors.HTTPError "Unauthorized"
// @Router /auth/login [post]
func Login(c echo.Context) error {
	req := new(AuthRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if !system.LL.State() {
		// check image captcha code
		captchaValue := system.CaptchaCode.Load()
		defer system.CaptchaCode.Store([]byte{})
		if captchaValue == nil {
			return errors.HTTPErrorCaptchaFailed
		}
		captchaCode := string(captchaValue.([]byte))
		if captchaCode == "" || captchaCode != req.Captcha {
			return errors.HTTPErrorCaptchaFailed
		}
	}

	if auth.Check(&req.AuthBody) == false {
		system.LL.Inc()
		return errors.HTTPErrorUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	t, e := token.SignedString([]byte(setting.Auth.SigningKey))
	if e != nil {
		return e
	}

	return c.JSON(http.StatusOK, &AuthResponse{t})
}

// @Summary deprecated
// @bolg captcha
// @Tags Auth
// @Accept  json
// @Produce  png
// @Success 200
// @Router /auth/captcha [get]
func Captcha(c echo.Context) error {
	f := system.NewCaptcha().Generate()
	return c.Stream(http.StatusOK, "image/png", f)
}

// todo

// @Summary deprecated
// @bolg system info
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/info [get]
func Info(c echo.Context) error {
	return nil
}
