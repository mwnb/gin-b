package main

import (
	"gin-b/options"
	"log"

	"github.com/gin-gonic/gin"
)

func run(engine *gin.Engine) {
	PORT := options.PORT
	err := engine.Run(PORT)
	if err != nil {
		log.Fatal("[run server fail err:]", err)
	}
}
