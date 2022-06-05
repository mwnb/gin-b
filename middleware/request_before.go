package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestBefore(ctx *gin.Context) {
	fmt.Println(ctx.Request.URL.Path, "<<<-QaQ")
}
