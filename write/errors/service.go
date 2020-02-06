package errors

const (
	ServiceResourceNotFoundErrorCode = 101
)

var (
	ServiceResourceNotFoundError = NewServiceError(ServiceResourceNotFoundErrorCode, "search resource not found")
)

var serviceCodeToHTTPErrorMap = map[int]*HTTPError{
	ServiceResourceNotFoundErrorCode: HTTPErrorNotFound,
}

type serviceError struct {
	code    int 	`json:"-"`
	Message string	`json:"message"`
}

func NewServiceError(code int, msg string) *serviceError {
	return &serviceError{
		code:    code,
		Message: msg,
	}
}

func (e *serviceError) Error() string {
	return e.Message
}

