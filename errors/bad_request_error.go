package errors

import (
	"fmt"
	"net/http"
)

type BadRequestError struct {
	ToserbaError
}

func NewBadRequestError(err string) *BadRequestError {
	return &BadRequestError{ToserbaError{message: err, code: http.StatusBadRequest}}
}

func NewBadRequestErrorf(format string, args ...interface{}) *BadRequestError {
	return &BadRequestError{ToserbaError{message: fmt.Sprintf(format, args...), code: http.StatusBadRequest}}
}
