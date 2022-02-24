package errors

const (
	ServiceResourceNotFoundErrorCode = 101
)

var (
	ServiceResourceNotFoundError = NewServiceError(ServiceResourceNotFoundErrorCode, "search resource not found")
)

var ServiceCodeToHTTPErrorMap = map[int]*HTTPError{
	ServiceResourceNotFoundErrorCode: HTTPErrorNotFound,
}

type ServiceError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewServiceError(code int, msg string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: msg,
	}
}

func (e *ServiceError) Error() string {
	return e.Message
}
