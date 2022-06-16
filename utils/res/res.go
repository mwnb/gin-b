package res

import "github.com/gin-gonic/gin"

func responseBase(flag bool) func(gin.H) gin.H {
	return func(j gin.H) gin.H {
		return gin.H{
			"success": flag,
			"data":    j,
		}
	}
}

var ResponseSuccess = responseBase(true)

var ResponseFail = responseBase(false)

func ResponseError(e error) gin.H {
	return gin.H{
		"success": false,
		"error":   e.Error(),
	}
}
