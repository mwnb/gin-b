package home

import (
	"fmt"
	"gin-b/common/dbinit"
	"gin-b/model"
	"gin-b/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomeImages(ctx *gin.Context) {
	fmt.Println("我是你爹")
	var images []model.Images
	var count int64
	dbinit.Db.Limit(15).Offset(0).Find(&images).Count(&count)
	fmt.Println(len(images), "images")
	ctx.JSON(http.StatusOK, res.ResponseSuccess(gin.H{
		"list":  images,
		"count": count,
	}))
}

func GetCarousel(ctx *gin.Context) {
	var carousels []model.Images
	dbinit.Db.Limit(2).Offset(0).Find(&carousels)
	ctx.JSON(http.StatusOK, res.ResponseSuccess(gin.H{
		"carousel": carousels,
	}))
}
