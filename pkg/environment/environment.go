package environment

import (
	"sync"
)

type environment struct {
	Host         string
	APIKeyID     string
	APIKeySecret string
}

var e *environment
var once sync.Once

/*GetInstance Export Singleton*/
func GetInstance(typeE typeEnvironment) *environment {

	once.Do(func() {
		e = typeE.getEnvironment()
	})

	return e
}
