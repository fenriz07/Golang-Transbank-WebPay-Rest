package client

import "fmt"

type errorClient struct {
	ErrorMessage string `json:"error_message"`
	Message      string
	HTTPCode     int
	HTTPStatus   string
}

func (e *errorClient) SetProperties(HTTPCode int, HTTPStatus string) {
	e.HTTPCode = HTTPCode
	e.HTTPStatus = HTTPStatus
	e.Message = fmt.Sprintf("Could not obtain a response from the service: %s (HTTP code %v)", HTTPStatus, HTTPCode)
}

func (e *errorClient) Error() string {
	return fmt.Sprintf(" Message: %s\n HTTPCode: %d\n HTTPStatus: %s\n ErrorMessage: %s ", e.Message, e.HTTPCode, e.HTTPStatus, e.ErrorMessage)
}
