package devgentoken

import (
	gentokenService "gin-b/service/dev/gen_token"

	"github.com/gin-gonic/gin"
)

const (
	Name = "DEV_GEN_TOKEN_CONTROLLER"
)

type devGenToken struct {
	Name string
}

func (d *devGenToken) GenToken() gin.HandlerFunc {
	return gentokenService.GenToken
}

var DevGenTokenController devGenToken = devGenToken{
	Name,
}
