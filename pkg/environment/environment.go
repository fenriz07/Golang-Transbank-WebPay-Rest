package environment

import (
	"sync"
)

/*Environment struct content data for client*/
type Environment struct {
	Host         string
	APIKeyID     string
	APIKeySecret string
}

var e *Environment
var once sync.Once

/*SetInstance set env instnce singleton*/
func SetInstance(env typeEnvironment) {

	once.Do(func() {
		e = env.getEnvironment()
	})
}

/*GetInstance get env instance*/
func GetInstance() *Environment {
	return e
}
