package main

import (
	"gin-b/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routers.MountRouter(app)
	run(app)
}
