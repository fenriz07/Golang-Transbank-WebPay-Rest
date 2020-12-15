package transaction

import (
	"testing"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
)

func TestCreate_success(t *testing.T) {

	env := environment.IntegrationEnviroment{}
	environment.SetInstance(env)
	client.SetInstance("webpay_success")

	resp, err := Create("test_success", "test_session_success", 1000, "http://test.com")

	if err != nil {
		t.Errorf("An error was not expected %v", err)
	}

	wantToken := "7452125458955"
	wantURL := "http://github.com"

	if resp.Token != wantToken && resp.URL != wantURL {
		t.Errorf("want token %v got %v", wantToken, resp.Token)
		t.Errorf("want url %v got %v", wantURL, resp.URL)
	}

}

func TestCommit_success(t *testing.T) {
	env := environment.IntegrationEnviroment{}
	environment.SetInstance(env)
	client.SetInstance("webpay_success")

	resp, err := Commit("token741456963")

	if err != nil {
		t.Errorf("An error was not expected %v", err)
	}

	if resp.ResponseCode != 0 {
		t.Errorf("want 0 got %v", resp.ResponseCode)
	}

}

func TestGetStatus(t *testing.T) {

	env := environment.IntegrationEnviroment{}
	environment.SetInstance(env)
	client.SetInstance("webpay_success")

	resp, err := GetStatus("getstatustokentest")

	if err != nil {
		t.Errorf("An error was not expected %v", err)
	}

	if resp.ResponseCode != 0 {
		t.Errorf("want 0 got %v", resp.ResponseCode)
	}

}

func TestRefund(t *testing.T) {
	env := environment.IntegrationEnviroment{}
	environment.SetInstance(env)
	client.SetInstance("webpay_success")

	resp, err := Refund("tokenrefundtest", 1000)

	if err != nil {
		t.Errorf("An error was not expected %v", err)
	}

	if resp.ResponseCode != 0 {
		t.Errorf("want 0 got %v", resp.ResponseCode)
	}

	if resp.NullifiedAmount != 1000 {
		t.Errorf("want 1000 got %v", resp.NullifiedAmount)
	}
}
