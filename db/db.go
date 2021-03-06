package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB
var err error
//初始化方法
func Init() {
	//链接数据库
	Db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/qiyun?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	err := Db.DB().Ping()
	if err != nil {
		Db.DB().Close()
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
}