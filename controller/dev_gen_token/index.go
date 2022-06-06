package devgentoken

import (
	devgentokenService "gin-b/service/dev_gen_token"

	"github.com/gin-gonic/gin"
)

const (
	Name = "DEV_GEN_TOKEN_CONTROLLER"
)

type devGenToken struct {
	Name string
}

func (d *devGenToken) GenToken() gin.HandlerFunc {
	return devgentokenService.GenToken
}

var DevGenTokenController devGenToken = devGenToken{
	Name,
}
