package errs

import (
	"fmt"
	"mimose/gcosy/magic"
)

type NError struct {
	Code int
	Desc string
	Err  error
}

// error message
func (e *NError) Error() string {
	errorMessage := "NError[code: %d]"
	if e.Desc != "" {
		errorMessage += " [desc: %s]"
	}
	if e.Err != nil {
		errorMessage += " [oriError: %s]"
	}
	return fmt.Sprintf(errorMessage, e.Code, e.Desc, magic.Ternary(e.Err != nil, e.Err.Error(), ""))
}

// new a NError
func New(code int, desc string, err error) *NError {
	return &NError{
		Code: code,
		Desc: desc,
		Err:  err,
	}
}
