package apierr

import (
	"net/http"
)

const InternalServerError = "Internal server error"

// Client is the APIERROR client
type Client struct {
}

// New return new client of api errors
func New() *Client {
	return &Client{}
}

// Message of response errors
type Message struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"error"`
	ErrorStatus  int    `json:"status"`
	Error        error  `json:"-"`
}

// BadRequest return new response in correct structure
func (c Client) BadRequest(message string, err error) *Message {
	return &Message{
		ErrorMessage: message,
		ErrorCode:    http.StatusText(http.StatusBadRequest),
		ErrorStatus:  http.StatusBadRequest,
		Error:        err,
	}
}

// InternalServerError return internal server error default
func (c Client) InternalServerError(err error) *Message {
	return &Message{
		ErrorMessage: InternalServerError,
		ErrorCode:    http.StatusText(http.StatusInternalServerError),
		ErrorStatus:  http.StatusInternalServerError,
		Error:        err,
	}
}
