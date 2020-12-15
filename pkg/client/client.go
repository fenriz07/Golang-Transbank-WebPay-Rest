package client

import (
	"sync"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
	"github.com/go-resty/resty/v2"
)

var w Client
var once sync.Once

/*SetInstance set instance singleton the HttpClient*/
func SetInstance(typeClient string) {

	env := environment.GetInstance()

	once.Do(func() {

		if typeClient == "http" {
			w = &HttpClient{Environment: env}
		} else if typeClient == "webpay_success" {
			w = &ClientWebPaySuccess{Environment: env}
		}
	})
}

/*GetInstance get env instance*/
func GetInstance() Client {
	return w
}

type Client interface {
	Get(endpoint string) (*resty.Response, error)
	Post(endpoint string, body map[string]interface{}) (*resty.Response, error)
	Put(endpoint string) (*resty.Response, error)
}
