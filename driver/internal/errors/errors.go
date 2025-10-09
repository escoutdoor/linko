package errors

import (
	"errors"
	"fmt"

	"github.com/escoutdoor/linko/driver/internal/errors/codes"
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

func DriverNotFoundWithID(driverID string) *Error {
	msg := fmt.Sprintf("driver with id %q was not found", driverID)
	return newError(codes.DriverNotFound, msg)
}

func ValidationFailed(msg string) *Error {
	return newError(codes.ValidationFailed, msg)
}
