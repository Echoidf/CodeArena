package common

type SysError struct {
	Code    int
	Message string
	Err     error
}

func (e *SysError) Error() string {
	return e.Message
}

var (
	InvalidInputError = &SysError{Code: 1001, Message: "please check the input data"}
	InsertError       = &SysError{Code: 1002, Message: "Insert error"}
)

func (e *SysError) WithError(err error) *SysError {
	e.Err = err
	return e
}
