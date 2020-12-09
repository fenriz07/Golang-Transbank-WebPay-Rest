package environment

const integrationHOST string = "https://webpay3gint.transbank.cl"
const integrationAPIKeyID string = "597055555532"
const integrationAPIKeySecret string = "579B532A7440BB0C9079DED94D31EA1615BACEB56610332264630D42D0A36B1C"

/*IntegrationEnviroment struct for test environement*/
type IntegrationEnviroment struct{}

func (e IntegrationEnviroment) getEnvironment() *environment {
	return &environment{
		Host:         integrationHOST,
		APIKeyID:     integrationAPIKeyID,
		APIKeySecret: integrationAPIKeySecret,
	}
}
