package main

import (
	"fmt"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
)

func main() {

	environmentIntegration := environment.IntegrationEnviroment{}

	env := environment.GetInstance(environmentIntegration)

	fmt.Println(env)

	env = environment.GetInstance(environmentIntegration)

	fmt.Println(env)
}
