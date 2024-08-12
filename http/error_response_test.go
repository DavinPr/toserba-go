package http_test

import (
	"testing"

	"github.com/DavinPr/toserba-go/http"
	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestError_WithoutMessage(t *testing.T) {
	resp := http.NewBadRequestError("")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "BAD_REQUEST",
				MessageTitle: "Bad Request",
				Message:      "can not process the request",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewBadRequestError_WithMessage(t *testing.T) {
	resp := http.NewBadRequestError("Test_Error")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "BAD_REQUEST",
				MessageTitle: "Bad Request",
				Message:      "Test_Error",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewBadRequestError_WithMultipleMessage(t *testing.T) {
	resp := http.NewBadRequestError("Test_Error1::Test_Error2")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "BAD_REQUEST",
				MessageTitle: "Bad Request",
				Message:      "Test_Error1",
			},
			{
				Code:         "BAD_REQUEST",
				MessageTitle: "Bad Request",
				Message:      "Test_Error2",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewNotFoundError(t *testing.T) {
	resp := http.NewNotFoundError()
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "NOT_FOUND",
				MessageTitle: "Not Found",
				Message:      "data not found",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewInternalServerError(t *testing.T) {
	resp := http.NewInternalServerError()
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "SERVER_ERROR",
				MessageTitle: "Internal Server Error",
				Message:      "Some error occurred, please try again",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewUnprocessableEntityError_WithoutMessage(t *testing.T) {
	resp := http.NewUnprocessableEntityError("")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "INVALID_REQUEST",
				MessageTitle: "Please Input Valid Request",
				Message:      "Request is invalid",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewUnprocessableEntityError_WithMessage(t *testing.T) {
	resp := http.NewUnprocessableEntityError("Test_Error")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "INVALID_REQUEST",
				MessageTitle: "Please Input Valid Request",
				Message:      "Test_Error",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewUnprocessableEntityError_WithMultipleMessage(t *testing.T) {
	resp := http.NewUnprocessableEntityError("Test_Error1::Test_Error2")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "INVALID_REQUEST",
				MessageTitle: "Please Input Valid Request",
				Message:      "Test_Error1",
			},
			{
				Code:         "INVALID_REQUEST",
				MessageTitle: "Please Input Valid Request",
				Message:      "Test_Error2",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewUnauthorizedError_WithoutMessage(t *testing.T) {
	resp := http.NewUnauthorizedError("")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "INVALID_AUTHORIZATION",
				MessageTitle: "You're not authorized",
				Message:      "You're not allowed to continue",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewUnauthorizedError_WithMessage(t *testing.T) {
	resp := http.NewUnauthorizedError("Test_Error")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "INVALID_AUTHORIZATION",
				MessageTitle: "You're not authorized",
				Message:      "Test_Error",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}

func TestNewUnauthorizedError_WithMultipleMessage(t *testing.T) {
	resp := http.NewUnauthorizedError("Test_Error1::Test_Error2")
	expectedResp := http.ServiceResponse{
		Success: false,
		Error: []http.ErrorDetail{
			{
				Code:         "INVALID_AUTHORIZATION",
				MessageTitle: "You're not authorized",
				Message:      "Test_Error1",
			},
			{
				Code:         "INVALID_AUTHORIZATION",
				MessageTitle: "You're not authorized",
				Message:      "Test_Error2",
			},
		},
	}
	assert.Equal(t, expectedResp, resp)
}
