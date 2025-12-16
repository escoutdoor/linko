package errors

import (
	"errors"
	"fmt"

	"github.com/escoutdoor/linko/catalog/internal/errors/codes"
)

type Error struct {
	Code codes.Code
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func newError(code codes.Code, err string) *Error {
	return &Error{
		Code: code,
		Err:  errors.New(err),
	}
}

func ValidationFailed(msg string) *Error {
	return newError(codes.ValidationFailed, msg)
}
