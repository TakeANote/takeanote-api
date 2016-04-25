package model

// Error describes an error that occured in the API with a message and HTTP status code
type Error struct {
	HTTPStatusCode int    `json:"status"`
	Message        string `json:"message"`
}

// NewError create a new Error struct that wrap error for JSON marshaling
func NewError(statusCode int, err error) Error {
	return Error{
		HTTPStatusCode: statusCode,
		Message:        err.Error(),
	}
}
