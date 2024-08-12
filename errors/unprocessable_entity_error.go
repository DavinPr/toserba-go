package errors

import (
	"fmt"
	"net/http"
)

type UnprocessableEntityError struct {
	ToserbaError
	title string
}

func (e *UnprocessableEntityError) ErrorTitle() string {
	return e.title
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		ToserbaError: ToserbaError{message: message, code: http.StatusUnprocessableEntity},
	}
}

func NewUnprocessableEntityErrorf(format string, args ...interface{}) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		ToserbaError: ToserbaError{message: fmt.Sprintf(format, args...), code: http.StatusUnprocessableEntity},
	}
}

func NewUnprocessableEntityErrorWithTitle(message string, title string) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		ToserbaError: ToserbaError{message: message, code: http.StatusUnprocessableEntity},
		title:        title,
	}
}
