package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ToserbaError
}

func (e *NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError(err string) *NotFoundError {
	return &NotFoundError{ToserbaError{message: err, code: http.StatusNotFound}}
}

func NewNotFoundErrorf(format string, args ...interface{}) *NotFoundError {
	return &NotFoundError{ToserbaError{message: fmt.Sprintf(format, args...), code: http.StatusNotFound}}
}
