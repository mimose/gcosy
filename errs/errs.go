package errs

import (
	"fmt"
	"mimose/gcosy/magic"
)

type CError struct {
	Code int
	Desc string
	Err  error
}

// error message
func (e *CError) Error() string {
	errorMessage := "CError[code: %d]"
	if e.Desc != "" {
		errorMessage += " [desc: %s]"
	}
	if e.Err != nil {
		errorMessage += " [oriError: %s]"
	}
	return fmt.Sprintf(errorMessage, e.Code, e.Desc, magic.Ternary(e.Err != nil, e.Err.Error(), ""))
}

// new a CError
func New(code int, desc string, err error) *CError {
	return &CError{
		Code: code,
		Desc: desc,
		Err:  err,
	}
}
