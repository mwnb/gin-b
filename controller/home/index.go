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

func (h *home) GetBanner() gin.HandlerFunc {
	return homeService.GetBanner
}

var HomeController home = home{
	Name,
}
