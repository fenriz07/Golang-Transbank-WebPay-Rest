package client

import (
	"sync"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
	"github.com/go-resty/resty/v2"
)

var w *Client
var once sync.Once

/*SetInstance set instance singleton the HttpClient*/
func SetInstance() {

	env := environment.GetInstance()

	once.Do(func() {
		ic := HttpClient{Environment: env}
		w = &ic
	})
}

/*GetInstance get env instance*/
func GetInstance() *HttpClient {
	return w
}

type Client interface {
	Get(endpoint string) (*resty.Response, error)
}
