package main

import (
	"fmt"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)

func main() {

	webpayplus.SetEnvironmentIntegration()
	//webpayplus.SetEnvironmentProduction("a","b")

	transaction, _ := transaction.Create("ordenCompra12345678", "sesion1234557545", 1000, "http://www.comercio.cl/webpay/retorno")

	fmt.Println(transaction)

}
