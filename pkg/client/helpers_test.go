package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
	"github.com/go-resty/resty/v2"
)

func TestGetHeaders(t *testing.T) {

	Host := "https://webpay3gint.transbank.cl/"
	APIKeyID := "597055555532"
	APIKeySecret := "579B532A7440BB0C9079DED94D31EA1615BACEB56610332264630D42D0A36B1C"

	want := map[string]string{
		"Tbk-Api-Key-Id":     APIKeyID,
		"Tbk-Api-Key-Secret": APIKeySecret,
		"Content-Type":       "application/json",
	}

	env := environment.Environment{
		Host:         Host,
		APIKeyID:     APIKeyID,
		APIKeySecret: APIKeySecret,
	}

	client := HttpClient{Environment: &env}

	got := client.getHeaders()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}

}

func TestGenerateURL(t *testing.T) {

	Host := "https://webpay3gint.transbank.cl/"
	APIKeyID := "597055555532"
	APIKeySecret := "579B532A7440BB0C9079DED94D31EA1615BACEB56610332264630D42D0A36B1C"

	env := environment.Environment{
		Host:         Host,
		APIKeyID:     APIKeyID,
		APIKeySecret: APIKeySecret,
	}

	client := HttpClient{Environment: &env}

	endpoint := "test"

	want := Host + endpoint
	got := client.generateURL(endpoint)

	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestValidateResponse_valid(t *testing.T) {

	resp := &resty.Response{
		RawResponse: &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"error_message":"json"}`))),
			Status:     "200 OK",
			StatusCode: 200,
		},
	}

	err := validateResponse(resp)

	if err != nil {
		t.Errorf("%v", err)
	}
}
