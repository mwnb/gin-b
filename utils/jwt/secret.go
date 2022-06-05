package jwt

import (
	"fmt"
	"gin-b/env"

	"gin-b/common"
)

var (
	secretKey = "b24a9920495e4e9da29e3265f6601be8f4a13579548f196f6dfcff0ab7fd2665833103e70daad7aec3d9d50408b7e233a73222b37891f079b73e7ed0067ad003"
)

func init() {
	if env.GIN_ENV == common.PRODUCTION {
		// 获取secretKey 给 变量 赋值
	} else {
		// 开发环境
		fmt.Printf("\033[1;37;42m%s\033[0m\n", common.DEVELOPMENT)
	}
}
