package errors

import "net/http"

const (
	httpDefaultErrorCode = http.StatusInternalServerError
	httpDefaultErrorKey  = "ServerError"
	httpDefaultErrorMsg  = ""
)

var (
	HTTPErrorUnauthorized = NewHTTPError(http.StatusUnauthorized, "Unauthorized", "认证失败")
	HTTPErrorUnauthenticated = NewHTTPError(http.StatusForbidden, "Unauthenticated", "未授权的访问")
	HTTPErrorNotFound = NewHTTPError(http.StatusNotFound, "NotFound", "资源不存在")
	HTTPErrorBadRequest = NewHTTPError(http.StatusBadRequest, "BadRequest", "参数错误")
)

type HTTPError struct {
	code    int 	`json:"-"`
	Key     string	`json:"key"`
	Message string	`json:"message"`
}

func NewHTTPError(code int, key string, msg string) *HTTPError {
	return &HTTPError{
		code:    code,
		Key:     key,
		Message: msg,
	}
}

func (e *HTTPError) Error() string {
	return e.Key + ": " + e.Message
}