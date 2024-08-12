package errors

import (
	"fmt"
	"net/http"
)

type UnauthorizedError struct {
	ToserbaError
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{ToserbaError{message: message, code: http.StatusUnauthorized}}
}

func NewUnauthorizedErrorf(format string, args ...interface{}) *UnauthorizedError {
	return &UnauthorizedError{ToserbaError{message: fmt.Sprintf(format, args...), code: http.StatusUnauthorized}}
}
