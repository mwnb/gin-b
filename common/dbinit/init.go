package dbinit

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	PORT    = 3306
	HOST    = "localhost"
	USER    = "root"
	PWD     = "123456"
	PROTOCL = "tcp"
	DB_NAME = "beautiful_girl_wallpaper"
)

func Conn() (*gorm.DB, error) {
	mysqlConInfo := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USER, PWD, PROTOCL, HOST, PORT, DB_NAME)
	return gorm.Open(mysql.Open(mysqlConInfo))
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = Conn()
	if err != nil {
		log.Fatal(err)
	}
}
