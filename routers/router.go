package routers

import (
	"gin-b/options"

	"github.com/gin-gonic/gin"

	"gin-b/controller/home"
)

func MountRouter(engine *gin.Engine) {
	router := engine.Group(options.PRE_FIX_PATH)

	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/getHomeImages", home.HomeController.GetHomeImages())
		homeRouter.GET("/getCarousel", home.HomeController.GetCarousel())
	}
}
