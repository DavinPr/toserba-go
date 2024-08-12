package errors

import (
	"fmt"
	"net/http"
)

type ConflictError struct {
	ToserbaError
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{ToserbaError{message: message, code: http.StatusConflict}}
}

func NewConflictErrorf(format string, args ...interface{}) *ConflictError {
	return &ConflictError{ToserbaError{message: fmt.Sprintf(format, args...), code: http.StatusConflict}}
}
