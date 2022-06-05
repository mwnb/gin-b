package env

import (
	"os"
	"strings"
)

var (
	GIN_ENV = "DEVELOPMENT"
)

func getGinEnv() {
	env := os.Environ()
	for _, v := range env {
		if strings.HasPrefix(v, "GIN_ENV") {
			GIN_ENV = strings.Split(v, "=")[1]
			break
		}
	}
}

func init() {
	getGinEnv()
}
