![Golang-WebPay](https://user-images.githubusercontent.com/9199380/102151890-a451d980-3e52-11eb-9626-3a679cce4850.png)

<p align="center">
<h1 align="center">Golang - Webpay Plus Transbank Rest</h1>
<p align="center">Biblioteca para la integración de Webpay Plus (REST) </p>

</p>

[![Go Report Card](https://goreportcard.com/badge/github.com/fenriz07/Golang-Transbank-WebPay-Rest)](https://goreportcard.com/report/github.com/fenriz07/Golang-Transbank-WebPay-Rest)

## Caracteristicas

- Soporte para ambiente de integración y producción
- Crear transacción.
- Confirmar una transacción.
- Obtener el estado de una transacción.
- Reversar o Anular un pago.
- Consolidación de respuestas en structs.
- Manejo de errores http.

## Instalación

```bash
go get github.com/fenriz07/Golang-Transbank-WebPay-Rest
```

## Uso

#### Inicializar ambiente

```go
//Importar el package webpayplus
import (
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)
```

##### Desarrollo

```go
/*
    Usar la función SetEnvironmentIntegration para el ambiente de desarrollo.
  Automaticamente configura las credenciales del comercio.
  Configura el cliente y todas las transacciones seran ejecutadas bajo este ambiente automaticamente

*/
webpayplus.SetEnvironmentIntegration()
```

##### Producción

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

## Creador

[Servio Z.](https://github.com/fenriz07)
