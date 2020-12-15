package main

import (
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)

func main() {

	webpayplus.SetEnvironmentIntegration()
	//webpayplus.SetEnvironmentProduction("a","b")

	transaction.Create("ordenCompra12345678", "sesion1234557545", 1000, "http://www.comercio.cl/webpay/retorno")

	//resp2, err := transaction.GetStatus("dakdawlkd")

	//fmt.Println(resp2)
	//fmt.Println(err)

}
