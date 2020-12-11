package environment

/*ProductionEnviroment struct for live environement*/
type ProductionEnviroment struct {
	Host         string
	APIKeyID     string
	APIKeySecret string
}

func (e ProductionEnviroment) getEnvironment() *Environment {
	return &Environment{
		Host:         e.Host,
		APIKeyID:     e.APIKeyID,
		APIKeySecret: e.APIKeySecret,
	}
}
