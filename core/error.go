package core

const (
	ErrCodeDefault     = 100
	ErrCodeLoginFailed = 101
)

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

func NewErr(message string, code int) *Err {
	return &Err{code, message}
}

func PanicErr(message string, code int) {
	panic(NewErr(message, code))
}

// P panic forever
func P(err error) {
	if err != nil {
		panic(err)
	}
}
