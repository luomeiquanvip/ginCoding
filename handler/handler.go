package handler

import (
	. "ginCoding/db"
	. "ginCoding/model"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func CreateAndInsert(c *gin.Context)  {
	//创建表
	Db.AutoMigrate(&Order{})
	//插入数据
	order1 := Order{Order_id: "111", User_name: "yi", Amount: 11.1,Status:"y",File_url:"www.baidu.com"}
	order2 := Order{Order_id: "222", User_name: "er", Amount: 22.2,Status:"n",File_url:"www.google.com"}
	order3 := Order{Order_id: "333", User_name: "san", Amount: 33.3,Status:"y",File_url:"www.yahu.com"}

	Db.Create(&order1)
	Db.Create(&order2)
	Db.Create(&order3)

	c.String(http.StatusOK, "create successful")
}
//获得所有订单信息
func GetOrdersApi(c *gin.Context) {
	var o *Order
	orders, err := o.GetOrders()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": orders,
	})

}
//更新操作
func UpdateOrder(c *gin.Context) {

	Db.Model(&Order{}).Where("amount = ?", 22.2).Update("User_name", "si")
	c.String(http.StatusOK, "update successful")
}

func FuzzySearch(c *gin.Context) {


	var o *Order
	orders, err := o.GetFuzzySearchs()
	if err !=nil {
		log.Fatal(err)
	}
	//H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK,gin.H{
		"result":orders,
		"count":len(orders),
	})

}