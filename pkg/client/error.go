package client

import "fmt"

type errorClient struct {
	ErrorMessage string `json:"error_message"`
	Message      string
	HTTPCode     int
	HTTPStatus   string
}

func (e *errorClient) Error() string {
	return fmt.Sprintf(" Message: %s\n HTTPCode: %d\n HTTPStatus: %s\n ErrorMessage: %s ", e.Message, e.HTTPCode, e.HTTPStatus, e.ErrorMessage)
}
