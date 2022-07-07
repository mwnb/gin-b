package home

import (
	homeService "gin-b/service/home"

	"github.com/gin-gonic/gin"
)

const (
	Name = "HOME_CONTROLLER"
)

type home struct {
	Name string
}

func (h *home) GetHomeImages() gin.HandlerFunc {
	return homeService.GetHomeImages
}

func (h *home) GetCarousel() gin.HandlerFunc {
	return homeService.GetCarousel
}

var HomeController home = home{
	Name,
}
