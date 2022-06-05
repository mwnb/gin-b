package main

import (
	"gin-b/middleware"
	"gin-b/routers"

	_ "gin-b/env"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	middleware.UseMiddleWare(app)
	routers.MountRouter(app)
	run(app)
}
