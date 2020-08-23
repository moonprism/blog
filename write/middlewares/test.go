package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)


func Test(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		time.Sleep(time.Duration(1) * time.Second)
		err = next(c)
		return
	}
}
