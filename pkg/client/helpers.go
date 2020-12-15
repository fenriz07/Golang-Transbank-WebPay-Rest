package client

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

func (client HttpClient) init() *resty.Client {
	clientResty := resty.New()

	headers := client.getHeaders()

	return clientResty.SetHeaders(headers)
}

func (client HttpClient) getHeaders() map[string]string {
	return map[string]string{
		"Tbk-Api-Key-Id":     client.Environment.APIKeyID,
		"Tbk-Api-Key-Secret": client.Environment.APIKeySecret,
		"Content-Type":       "application/json",
	}
}

func (client HttpClient) generateURL(endpoint string) string {
	return client.Environment.Host + endpoint
}

func validateResponse(resp *resty.Response) error {

	var httpCode int = resp.StatusCode()
	var httpStatus string = resp.Status()

	if httpCode != 200 && httpCode != 204 {

		var errClient errorClient

		errClient.SetProperties(httpCode, httpStatus)

		err := json.Unmarshal(resp.Body(), &errClient)

		if err != nil {
			log.Printf("Error Unmarshal response server :  %v\n", err)
		}

		return &errClient
	}

	return nil
}
