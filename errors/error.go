package errors

import (
	goErrors "errors"
	pkgErrors "github.com/pkg/errors"
	"net/http"
)

type baseError interface {
	Error() string
	Code() int
}

type ToserbaError struct {
	message string
	code    int
}

func (e ToserbaError) Error() string {
	return e.message
}

func (e ToserbaError) Code() int {
	if e.code != 0 {
		return e.code
	}
	return http.StatusInternalServerError
}

func NewToserbaError(message string, code int) ToserbaError {
	return ToserbaError{message: message, code: code}
}

func New(message string) ToserbaError {
	return ToserbaError{message: message}
}

func Wrap(err error, message string) error {
	return pkgErrors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return pkgErrors.Wrapf(err, format, args...)
}

func ValidateErrorType[T baseError](err error) bool {
	var target T
	return goErrors.As(err, &target)
}
