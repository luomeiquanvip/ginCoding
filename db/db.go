package db

import (
	"fmt"
	"ginCoding/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

//初始化方法
func Init() {
	//链接数据库
	db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/qiyun?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//创建表
	var Order *model.Order
	db.CreateTable(&Order)

	db.AutoMigrate(&model.Order{})
}
