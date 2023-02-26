package model

// ErrorResponse will be used to wrap errors into a JSON response to send back to client
type ErrorResponse struct {
	Message string `json:"message"`
}
