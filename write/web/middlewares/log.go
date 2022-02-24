package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func SetLog(l *logrus.Logger) {
	logger = l
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.WithFields(logrus.Fields{
			"method": c.Request().Method,
			"uri":    c.Request().URL.String(),
			"ip":     c.Request().RemoteAddr,
		}).Info("incoming request")
		return next(c)
	}
}
