package environment

/*ProductionEnviroment struct for live environement*/
type ProductionEnviroment struct {
	Host         string
	APIKeyID     string
	APIKeySecret string
}

func (e ProductionEnviroment) getEnvironment() *environment {
	return &environment{
		Host:         e.Host,
		APIKeyID:     e.APIKeyID,
		APIKeySecret: e.APIKeySecret,
	}
}
