package env

import (
	"fmt"
	"gin-b/common"
	"os"
)

var (
	GIN_ENV = "DEVELOPMENT"
)

func getGinEnv() {
	GIN_ENV = os.Getenv("GIN_ENV")
}

func init() {
	getGinEnv()
	if GIN_ENV == common.PRODUCTION {
		// 获取secretKey 给 变量 赋值
	} else if GIN_ENV == common.DEVELOPMENT {
		// 开发环境
		fmt.Printf("\033[1;37;42m%s\033[0m\n", common.DEVELOPMENT)
	} else {
		panic(">>> 请用.sh启动! <<<")
	}
}
