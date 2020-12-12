package client

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
	"github.com/go-resty/resty/v2"
)

var w *HttpClient
var once sync.Once

/*HttpClient client*/
type HttpClient struct {
	Environment *environment.Environment
}

/*SetInstance set instance singleton the HttpClient*/
func SetInstance() {

	env := environment.GetInstance()

	once.Do(func() {
		w = &HttpClient{Environment: env}
	})
}

/*GetInstance get env instance*/
func GetInstance() *HttpClient {
	return w
}

/*Post http post for client*/
func (client HttpClient) Post(endpoint string, body map[string]interface{}) (*resty.Response, error) {
	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		SetBody(body).
		Post(url)

	if err != nil {
		log.Printf("Resty fail in post method client :  %v\n", err)
		return resp, err
	}

	err = client.validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err
}

/*Put http put method for client*/
func (client HttpClient) Put(endpoint string) (*resty.Response, error) {
	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		Put(url)

	if err != nil {
		log.Printf("Resty fail in post method client :  %v\n", err)
		return resp, err
	}

	err = client.validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err
}

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

func (client HttpClient) validateResponse(resp *resty.Response) error {

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
