package errors

import "net/http"

const (
	HttpDefaultErrorCode = http.StatusInternalServerError
	HttpDefaultErrorKey  = "ServerError"
	HttpDefaultErrorMsg  = ""
)

var (
	HTTPErrorUnauthorized    = NewHTTPError(http.StatusUnauthorized, "Unauthorized", "认证失败")
	HTTPErrorCaptchaFailed    = NewHTTPError(http.StatusUnauthorized, "CaptchaFailed", "验证码错误")
	HTTPErrorUnauthenticated = NewHTTPError(http.StatusForbidden, "Unauthenticated", "未授权的访问")
	HTTPErrorNotFound        = NewHTTPError(http.StatusNotFound, "NotFound", "资源不存在")
	HTTPErrorBadRequest      = NewHTTPError(http.StatusBadRequest, "BadRequest", "参数错误")
)

type HTTPError struct {
	Code    int    `json:"-"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

func NewHTTPError(code int, key string, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Key:     key,
		Message: msg,
	}
}

func (e *HTTPError) Error() string {
	return e.Key + ": " + e.Message
}
