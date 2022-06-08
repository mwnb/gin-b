package gentoken

import (
	"gin-b/utils/jwt"
	"gin-b/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenToken(ctx *gin.Context) {
	name, account, extra := ctx.Query("name"), ctx.Query("account"), ctx.Query("extra")
	token, err := jwt.CreateToken(name, account, extra)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res.ResponseFail(gin.H{
			"message": "gen token fail",
			"token":   "",
		}))
	} else {
		ctx.JSON(http.StatusOK, res.ResponseSuccess(gin.H{
			"message": "gen token success",
			"token":   token,
		}))
	}
}
