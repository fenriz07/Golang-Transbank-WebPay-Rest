package main

import (
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
)

func main() {

	environmentIntegration := environment.IntegrationEnviroment{}

	env := environment.GetInstance(environmentIntegration)

	webPayClient := client.WebPayClient{Environment: env}

	webPayClient.Create("ordenCompra12345678", "sesion1234557545", 1000, "http://www.comercio.cl/webpay/retorno")
}
