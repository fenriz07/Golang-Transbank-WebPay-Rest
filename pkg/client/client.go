package client

import (
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

/*Get http het method for client*/
func (client HttpClient) Get(endpoint string) (*resty.Response, error) {

	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		Get(url)

	if err != nil {
		log.Printf("Resty fail in get method client :  %v\n", err)
		return resp, err
	}

	err = client.validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err

}
