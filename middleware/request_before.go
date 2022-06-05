package middleware

import (
	_ "gin-b/utils/jwt"
	"gin-b/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func requestBefore(ctx *gin.Context) {
	// 验证 token
	if true {

	} else {
		ctx.JSON(http.StatusForbidden, res.ResponseFail(gin.H{
			"auth": false,
		}))
		ctx.Abort()
	}
}
