package errors

import "fmt"

type DBError struct {
	ToserbaError
}

func NewSimpleDBError(message string) *DBError {
	return &DBError{ToserbaError{message: message}}
}

func NewSimpleDBErrorf(format string, args ...interface{}) *DBError {
	return &DBError{ToserbaError{message: fmt.Sprintf(format, args...)}}
}

func NewDBError(err error, message string) *DBError {
	return &DBError{ToserbaError{message: fmt.Sprintf("%s: %s", message, err.Error())}}
}

func NewDBErrorf(err error, format string, args ...interface{}) *DBError {
	message := fmt.Sprintf(format, args...)
	return &DBError{ToserbaError{message: fmt.Sprintf("%s: %s", message, err.Error())}}
}
