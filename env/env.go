package env

import (
	"os"
)

var (
	GIN_ENV = "DEVELOPMENT"
)

func getGinEnv() {
	GIN_ENV = os.Getenv("GIN_ENV")
}

func init() {
	getGinEnv()
}
