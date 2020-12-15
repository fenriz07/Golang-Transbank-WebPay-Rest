package client

import (
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

var clientResty *resty.Client = resty.New()
var fakeURL string = "https://api.mybiz.com/articles.json"

type ClientWebPaySuccess struct {
	Environment *environment.Environment
}

/*Post http post for client*/
func (client ClientWebPaySuccess) Post(endpoint string, body map[string]interface{}) (*resty.Response, error) {

	var fixture string = ""

	if endpoint == "rswebpaytransaction/api/webpay/v1.0/transactions" {
		fixture = `{"token":"7452125458955","url":"http://github.com"}`
	} else if endpoint == "rswebpaytransaction/api/webpay/v1.0/transactions/tokenrefundtest/refunds" {
		fixture = `{
			"type":"NULLIFY",
			"authorization_code":"123456",
			"authorization_date":"2019-03-20T20:18:20Z",
			"nullified_amount":1000.00,
			"balance":0.00,
			"response_code":0
		 }`
	}

	responder := httpmock.NewStringResponder(200, fixture)
	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(clientResty.GetClient())

	httpmock.RegisterResponder("POST", fakeURL, responder)

	resp, err := clientResty.R().Post(fakeURL)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*Put http put method for client*/
func (client ClientWebPaySuccess) Put(endpoint string) (*resty.Response, error) {

	fixture := `{
		"vci": "TSY",
		"amount": 10000,
		"status": "AUTHORIZED",
		"buy_order": "ordenCompra12345678",
		"session_id": "sesion1234557545",
		"card_detail": {
			"card_number": "6623"
		},
		"accounting_date": "0522",
		"transaction_date": "2019-05-22T16:41:21.063Z",
		"authorization_code": "1213",
		"payment_type_code": "VN",
		"response_code": 0,
		"installments_number": 0
	  }`

	responder := httpmock.NewStringResponder(200, fixture)

	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(clientResty.GetClient())

	httpmock.RegisterResponder("PUT", fakeURL, responder)

	resp, err := clientResty.R().Put(fakeURL)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*Get http het method for client*/
func (client ClientWebPaySuccess) Get(endpoint string) (*resty.Response, error) {

	fixture := `{
		"vci": "TSY",
		"amount": 10000,
		"status": "AUTHORIZED",
		"buy_order": "ordenCompra12345678",
		"session_id": "sesion1234557545",
		"card_detail": {
			"card_number": "6623"
		},
		"accounting_date": "0522",
		"transaction_date": "2019-05-22T16:41:21.063Z",
		"authorization_code": "1213",
		"payment_type_code": "VN",
		"response_code": 0,
		"installments_number": 0
	  }`

	responder := httpmock.NewStringResponder(200, fixture)

	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(clientResty.GetClient())

	httpmock.RegisterResponder("GET", fakeURL, responder)

	resp, err := clientResty.R().Get(fakeURL)

	if err != nil {
		return resp, err
	}

	return resp, nil

}
