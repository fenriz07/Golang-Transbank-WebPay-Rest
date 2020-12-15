package environment

const productionHOST string = "https://webpay3g.transbank.cl/"

/*ProductionEnviroment struct for live environement*/
type ProductionEnviroment struct {
	Host         string
	APIKeyID     string
	APIKeySecret string
}

func (e ProductionEnviroment) getEnvironment() *Environment {
	return &Environment{
		Host:         productionHOST,
		APIKeyID:     e.APIKeyID,
		APIKeySecret: e.APIKeySecret,
	}
}

/*SetCredentials func set credentiales for production*/
func (e ProductionEnviroment) SetCredentials(APIKeyID string, APIKeySecret string) {
	e.APIKeyID = APIKeyID
	e.APIKeySecret = APIKeySecret
}
