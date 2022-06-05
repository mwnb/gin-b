package middleware

import (
	"github.com/gin-gonic/gin"
)

func requestBefore(ctx *gin.Context) {
	// 验证 token
	if true {

	} else {
		ctx.Abort()
	}
}
