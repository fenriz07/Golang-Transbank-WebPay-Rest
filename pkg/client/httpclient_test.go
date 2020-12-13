package client

import (
	"testing"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
)

func TestPost_response_invalid(t *testing.T) {

	env := environment.Environment{
		Host:         "https://jsonplaceholder.typicode.com/",
		APIKeyID:     "test-api-key",
		APIKeySecret: "test-api-secret",
	}

	httpClient := HttpClient{Environment: &env}

	resp, err := httpClient.Post("posts", map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	})

	if err == nil {
		t.Errorf("want 201 got %d", resp.StatusCode())
	}
}

func TestPost_response_validate(t *testing.T) {

	env := environment.Environment{
		Host:         "https://webpay3gint.transbank.cl/",
		APIKeyID:     "597055555532",
		APIKeySecret: "579B532A7440BB0C9079DED94D31EA1615BACEB56610332264630D42D0A36B1C",
	}

	httpClient := HttpClient{Environment: &env}

	body := map[string]interface{}{
		"buy_order":  "test-golang-transbank-1",
		"session_id": "hola",
		"amount":     1000,
		"return_url": "https://google.com/",
	}

	resp, err := httpClient.Post("rswebpaytransaction/api/webpay/v1.0/transactions", body)

	if err != nil {
		t.Errorf("want not error %v", err)
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 204 {
		t.Errorf("want 200 or 204, got %d", resp.StatusCode())
	}
}

func TestPut_response_invalid(t *testing.T) {

	env := environment.Environment{
		Host:         "https://jsonplaceholder.typicode.com/",
		APIKeyID:     "test-api-key",
		APIKeySecret: "test-api-secret",
	}

	httpClient := HttpClient{Environment: &env}

	_, err := httpClient.Put("posts/1")

	if err != nil {
		t.Errorf("got error in method put %v", err)
	}
}

func TestGet_response_valid(t *testing.T) {
	env := environment.Environment{
		Host:         "https://jsonplaceholder.typicode.com/",
		APIKeyID:     "test-api-key",
		APIKeySecret: "test-api-secret",
	}

	httpClient := HttpClient{Environment: &env}

	_, err := httpClient.Get("posts")

	if err != nil {
		t.Errorf("got error in method get %v", err)
	}
}
