package http

type ErrorDetail struct {
	Code         string `json:"code"`
	Message      string `json:"message"`
	MessageTitle string `json:"message_title"`
}

type ServiceResponse struct {
	Success bool          `json:"success"`
	Data    interface{}   `json:"data,omitempty"`
	Error   []ErrorDetail `json:"errors,omitempty"`
}

func NewSuccessServiceResponse(data interface{}) *ServiceResponse {
	return &ServiceResponse{
		Success: true,
		Data:    data,
	}
}
