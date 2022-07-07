package middleware

import (
	_ "gin-b/utils/jwt"

	"github.com/gin-gonic/gin"
)

func requestBefore(ctx *gin.Context) {
	// 验证 token
	// token := ctx.Request.Header["Token"]
	// if len(token) == 0 {
	// 	token = append(token, "")
	// }
	// jd, err := jwt.ParseToken(token[0])
	// if err != nil {
	// 	ctx.JSON(http.StatusForbidden, res.ResponseFail(gin.H{
	// 		"auth": false,
	// 	}))
	// 	ctx.Abort()
	// } else {
	// 	ctx.Set("username", jd.Name)
	// 	ctx.Set("account", jd.Account)
	// 	ctx.Set("extra", jd.Extra)
	// }
}
