package apierr

import (
	"net/http"
)

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
	Error        error
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
