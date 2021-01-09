![Golang-WebPay](https://user-images.githubusercontent.com/9199380/102151890-a451d980-3e52-11eb-9626-3a679cce4850.png)

<p align="center">
<h1 align="center">Golang - Webpay Plus Transbank Rest</h1>
<p align="center">Biblioteca para la integración de Webpay Plus (REST) </p>

</p>

<p align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/fenriz07/Golang-Transbank-WebPay-Rest)](https://goreportcard.com/report/github.com/fenriz07/Golang-Transbank-WebPay-Rest)
![Test](https://github.com/fenriz07/Golang-Transbank-WebPay-Rest/workflows/Test/badge.svg)
<a href="https://pkg.go.dev/github.com/fenriz07/Golang-Transbank-WebPay-Rest"><img src="https://godoc.org/github.com/fenriz07/Golang-Transbank-WebPay-Rest?status.svg" alt="GoDoc"></a>

</p>

## Características

- Soporte para ambiente de integración y producción
- Crear transacción.
- Confirmar una transacción.
- Obtener el estado de una transacción.
- Reversar o Anular un pago.
- Consolidación de respuestas en structs.
- Manejo de errores http.

## Ejemplo de uso (Gorilla):

Ejemplo de proyecto donde esta implementado
https://github.com/fenriz07/EjemploTransbankGoRest

## Instalación

```bash
go get github.com/fenriz07/Golang-Transbank-WebPay-Rest
```

## Uso

### Inicializar ambiente

Hay 2 ambientes Integración y producción

```go
//Importar el package webpayplus
import (
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)
```

#### Integración

```go
/*
    Usar la función SetEnvironmentIntegration para el ambiente de desarrollo.
  Automaticamente configura las credenciales del comercio.
  Configura el cliente y todas las transacciones seran ejecutadas bajo este ambiente automaticamente

*/
webpayplus.SetEnvironmentIntegration()
```

#### Producción

```go
/*
      Usar la función SetEnvironmentProduction para el ambiente de desarrollo.
  Automaticamente configura las credenciales del comercio.
  Configura el cliente y todas las transacciones seran ejecutadas bajo este ambiente automaticamente

  Dicha función acepta 2 parametros
  1 - APIKeyID     (string)
  2 - APIKeySecret (string)
*/

APIKeyID := "Código de comercio"
APIKeySecret := "Llave secreta"

webpayplus.SetEnvironmentProduction(APIKeyID,APIKeySecret)
```

### Crear una transacción Webpay Plus

Para crear una transacción basta llamar al método `transaction.Create()`

```go
/* Se deben pasar los siguentes parametros
  buyOrder = numero de la orden (string)
  sessionID = session (string)
  amount = monto a cobrar (int)
  returnURL = url de retorno que devolvera webpay
*/
Create(buyOrder string, sessionID string, amount int, returnURL string)
```

El mismo retornara el siguente `struct` de respuesta:

```go
/*TransactionCreateResponse struct with contain skeleton json to createResponse*/
type TransactionCreateResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}
```

### Confirmar una transacción Webpay Plus

Cuando el comercio retoma el control mediante return_url debes confirmar y obtener el resultado de una transacción usando el método `transaction.Commit()`

```go
/* Se deben pasar los siguentes parametros
  token = token devuelto por transbank  (string)
*/
Commit(token string) (response.TransactionCommitResponse, error)
```

El mismo retornara el siguente `struct` de respuesta:

```go
/*TransactionCommitResponse struct with contain skeleton json to commitResponse*/
type TransactionCommitResponse struct {
	Vci        string `json:"vci"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	BuyOrder   string `json:"buy_order"`
	SessionID  string `json:"session_id"`
	CardDetail struct {
		CardNumber string `json:"card_number"`
	} `json:"card_detail"`
	AccountingDate     string    `json:"accounting_date"`
	TransactionDate    time.Time `json:"transaction_date"`
	AuthorizationCode  string    `json:"authorization_code"`
	PaymentTypeCode    string    `json:"payment_type_code"`
	ResponseCode       int       `json:"response_code"`
	InstallmentsNumber int       `json:"installments_number"`
}

```

### Obtener estado de una transacción Webpay Plus

Esta operación permite obtener el estado de la transacción en cualquier momento. En condiciones normales es probable que no se requiera ejecutar, pero en caso de ocurrir un error inesperado permite conocer el estado y tomar las acciones que correspondan. `transaction.GetStatus()`

```go
/* Se deben pasar los siguentes parametros
  token = token devuelto por transbank  (string)
*/
GetStatus(token string) (response.TransactionStatusResponse, error)
```

El mismo retornara el siguente `struct` de respuesta:

```go
/*TransactionCommitResponse struct with contain skeleton json to commitResponse*/
type TransactionCommitResponse struct {
	Vci        string `json:"vci"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	BuyOrder   string `json:"buy_order"`
	SessionID  string `json:"session_id"`
	CardDetail struct {
		CardNumber string `json:"card_number"`
	} `json:"card_detail"`
	AccountingDate     string    `json:"accounting_date"`
	TransactionDate    time.Time `json:"transaction_date"`
	AuthorizationCode  string    `json:"authorization_code"`
	PaymentTypeCode    string    `json:"payment_type_code"`
	ResponseCode       int       `json:"response_code"`
	InstallmentsNumber int       `json:"installments_number"`
}

```

### Reversar o Anular un pago Webpay Plus

Este método permite a todo comercio habilitado, reembolsar o anular una transacción que fue generada en Webpay Plus. `transaction.Refund()`

```go
/* Se deben pasar los siguentes parametros
  token = token devuelto por transbank  (string)
  amount = monto a reembolsar o anular (int)
*/
func Refund(token string, amount int) (response.TransactionRefundResponse, error)
```

El mismo retornara el siguente `struct` de respuesta:

```go
/*TransactionRefundResponse struct with contain skeleton json to refundResponse*/
type TransactionRefundResponse struct {
	Type              string    `json:"type"`
	AuthorizationCode string    `json:"authorization_code"`
	AuthorizationDate time.Time `json:"authorization_date"`
	NullifiedAmount   float64   `json:"nullified_amount"`
	Balance           float64   `json:"balance"`
	ResponseCode      int       `json:"response_code"`
}

```

## Otras bibliotecas para los servicios de transbank:
- [OneClick Mall](https://github.com/fenriz07/Golang-Transbank-Oneclick-mall)

## Creador

[Fenriz07](https://github.com/fenriz07)
