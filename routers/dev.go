package routers

import (
	"gin-b/common"
	devgentoken "gin-b/controller/dev_gen_token"
	"gin-b/env"
	"gin-b/options"

	"github.com/gin-gonic/gin"
)

func MountDevRouter(engine *gin.Engine) {
	router := engine.Group(options.PRE_FIX_PATH)

	if env.GIN_ENV == common.DEVELOPMENT {
		devRouter := router.Group("/dev")
		devRouter.GET("/genToken", devgentoken.DevGenTokenController.GenToken())
	}
}
