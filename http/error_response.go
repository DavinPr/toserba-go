package http

import (
	"strings"
)

func NewBadRequestError(message string) ServiceResponse {
	if message == "" {
		message = "can not process the request"
	}
	return wrapErrorResponse(getErrorDetails("BAD_REQUEST", "Bad Request", message))
}

func NewNotFoundError() ServiceResponse {
	return wrapErrorResponse(getErrorDetails("NOT_FOUND", "Not Found", "data not found"))
}

func NewInternalServerError() ServiceResponse {
	return wrapErrorResponse(getErrorDetails("SERVER_ERROR", "Internal Server Error", "Some error occurred, please try again"))
}

func NewUnprocessableEntityError(message string) ServiceResponse {
	if message == "" {
		message = "Request is invalid"
	}
	return wrapErrorResponse(getErrorDetails("INVALID_REQUEST", "Please Input Valid Request", message))
}

func NewUnauthorizedError(message string) ServiceResponse {
	if message == "" {
		message = "You're not allowed to continue"
	}
	return wrapErrorResponse(getErrorDetails("INVALID_AUTHORIZATION", "You're not authorized", message))
}

func wrapErrorResponse(detail []ErrorDetail) ServiceResponse {
	return ServiceResponse{
		Success: false,
		Error:   detail,
	}
}

func getErrorDetails(code, messageTitle, msgStr string) []ErrorDetail {
	detail := ErrorDetail{Code: code, MessageTitle: messageTitle, Message: msgStr}
	messages := strings.Split(msgStr, "::")

	details := make([]ErrorDetail, 0)
	for _, msg := range messages {
		detail.Message = msg
		details = append(details, detail)
	}

	return details
}
