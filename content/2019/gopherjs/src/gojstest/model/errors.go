package model

import "fmt"

// ErrInvalid means a validation failure.
type ErrInvalid string

func ErrInvalidF(format string, a ...interface{}) ErrInvalid {
	return ErrInvalid(fmt.Sprintf(format, a))
}

// Error describes the detail.
func (e ErrInvalid) Error() string {
	return string(e)
}
