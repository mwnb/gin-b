package home

import (
	"gin-b/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBanner(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, res.ResponseSuccess(gin.H{
		"list": []string{},
	}))
}
