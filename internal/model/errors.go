package model

type HttpError struct {
}

func (h *HttpError) Error() string {
	return "http request error"
}

type JsonError struct {
}

func (j *JsonError) Error() string {
	return "json deserialization error during decoding response"
}

type InvalidRepositoryType struct {
}

func (i *InvalidRepositoryType) Error() string {
	return "initializing repo failed, please check REPOSITORY_TYPE environment variable"
}

// ErrorResponse will be used to wrap errors into a JSON response to send back to client
type ErrorResponse struct {
	Message string `json:"message"`
}
