package main

import (
	"fmt"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction"
)

func main() {

	environmentIntegration := environment.IntegrationEnviroment{}

	environment.SetInstance(environmentIntegration)

	client.SetInstance()

	transaction, _ := transaction.Create("ordenCompra12345678", "sesion1234557545", 1000, "http://www.comercio.cl/webpay/retorno")

	fmt.Println(transaction.URL)
}
