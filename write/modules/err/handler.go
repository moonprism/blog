package err

import (
	"git.kicoe.com/blog/write/modules/err/errors"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var httpErr = errors.NewHTTPError(errors.HttpDefaultErrorCode, errors.HttpDefaultErrorKey, errors.HttpDefaultErrorMsg)

	if he, ok := err.(*errors.HTTPError); ok {
		httpErr = he
	} else if se, ok := err.(*errors.ServiceError); ok {
		if he, has := errors.ServiceCodeToHTTPErrorMap[se.Code]; has {
			httpErr = he
		} else {
			// todo debug custom service error
		}
	} else {
		// todo debug
		httpErr.Message = err.Error()
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(httpErr.Code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(httpErr.Code, httpErr)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

func CustomJwtErrorHandler(err error) error {
	return errors.HTTPErrorUnauthenticated
}
