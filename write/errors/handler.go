package errors

import (
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var httpErr = NewHTTPError(httpDefaultErrorCode, httpDefaultErrorKey, httpDefaultErrorMsg)

	if he, ok := err.(*HTTPError); ok {
		httpErr = he
	} else if se, ok := err.(*serviceError); ok {
		if he, has := serviceCodeToHTTPErrorMap[se.code]; has {
			httpErr = he
		} else {
			// debug custom service error
		}
	} else {
    	// debug
        httpErr.Message = err.Error()
    }

    if !c.Response().Committed {
        if c.Request().Method == echo.HEAD {
            err := c.NoContent(httpErr.code)
            if err != nil {
                c.Logger().Error(err)
            }
        } else {
            err := c.JSON(httpErr.code, httpErr)
            if err != nil {
                c.Logger().Error(err)
            }
        }
    }
}

func CustomJwtErrorHandler(err error) error {
	return HTTPErrorUnauthenticated
}