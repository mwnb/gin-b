package middleware

import "github.com/gin-gonic/gin"

func UseMiddleWare(ctx *gin.Engine) {
	ctx.Use(requestBefore)
}
