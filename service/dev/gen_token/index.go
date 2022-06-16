package gentoken

import (
	"gin-b/utils/jwt"
	"gin-b/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	http://localhost:7878/gin/api/dev/genToken?name={}&account={}
	copy上面的地址在浏览器中执行
*/
func GenToken(ctx *gin.Context) {
	var gtp GenTokenParameter
	if err := ctx.ShouldBind(&gtp); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ResponseError(err))
		return
	}
	if gtp.Extra == "" {
		gtp.Extra = "{}"
	}
	token, err := jwt.CreateToken(gtp.Name, gtp.Account, gtp.Extra)
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
