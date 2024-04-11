package core

const (
	CodeDefault     = 0
	CodeLoginFailed = 100
)

// PfError Custom Error
type PfError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (pe *PfError) Error() string {
	return pe.Message
}

func NewPfError(message string, code int) PfError {
	return PfError{code, message}
}

// Pf panic forever
func Pf(err error, codes ...int) {
	if err != nil {
		code := CodeDefault
		if len(codes) > 0 {
			code = codes[0]
		}
		panic(NewPfError(err.Error(), code))
	}
}
