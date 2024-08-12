package errors_test

import (
	"testing"

	"github.com/DavinPr/toserba-go/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidateErrorType_True(t *testing.T) {
	err := errors.NewNotFoundErrorf("test error")

	assert.Error(t, err)
	assert.True(t, errors.ValidateErrorType[*errors.NotFoundError](err))
}

func TestValidateErrorType_False(t *testing.T) {
	err := errors.New("test error")

	assert.Error(t, err)
	assert.False(t, errors.ValidateErrorType[*errors.NotFoundError](err))
}
