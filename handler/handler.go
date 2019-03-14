package handler

import (
	"fmt"
	. "ginCoding/db"
	. "ginCoding/model"
	"github.com/mapslice"
	"github.com/tealeg/xlsx"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"log"
	"net/http"
	"os"
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
func Orders(c *gin.Context) {
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
//模糊查询
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

func InfoOrderByAmount(c *gin.Context) {
	var o *Order
	orders, err := o.GetInfoOrderByAmount()
	if err !=nil {
		log.Fatal(err)
	}
	//H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK,gin.H{
		"result":orders,
		"count":len(orders),
	})
}

func TxInsert(c *gin.Context)  {


	var o *Order
	orders, err := o.TxInsert()
	if err !=nil {
		log.Fatal(err)
	}
	//H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK,gin.H{
		"result":orders,
		"count":len(orders),
	})
}

func UploadFile(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename
	fmt.Println(file, err, filename)

	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful \n")
}




func DownloadToExcel(c *gin.Context) {
	var order OrderBack
	var orderlists []OrderBack
	orderlist, err := GetorderList()
	if err !=nil{
		log.Fatal(err)
	}

	for i := 0; i < len(orderlist); i++ {
		order.ID =  orderlist[i].ID
		order.Order_id = orderlist[i].Order_id
		order.User_name = orderlist[i].User_name
		order.Amount =  orderlist[i].Amount
		order.Status = orderlist[i].Status
		order.File_url = orderlist[i].File_url
		orderlists = append(orderlists, order)
	}
	id, _ := mapslice.ToStrings(orderlists, "ID")
	order_id, _ := mapslice.ToStrings(orderlists, "Order_id")
	user_name, _ := mapslice.ToStrings(orderlists, "User_name")
	amount, _ := mapslice.ToStrings(orderlists, "Amount")
	status, _ := mapslice.ToStrings(orderlists, "Status")
	file_url, _ := mapslice.ToStrings(orderlists, "File_url")


	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "编号"
	cell = row.AddCell()
	cell.Value = "订单编号"
	cell = row.AddCell()
	cell.Value = "用户名"
	cell = row.AddCell()
	cell.Value = "状态"
	cell = row.AddCell()
	cell.Value = "订单路径"
	for i := 0; i < len(id); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = id[i]
		cell = row.AddCell()
		cell.Value = order_id[i]
		cell = row.AddCell()
		cell.Value = user_name[i]
		cell = row.AddCell()
		cell.Value = amount[i]
		cell = row.AddCell()
		cell.Value = status[i]
		cell = row.AddCell()
		cell.Value = file_url[i]
		file.Save("/home/qydev/go/src/ginCoding/File.xlsx")
	}

	c.String(http.StatusCreated, "DownloadToExcel successful \n")
}
func GetorderList() (orderlist []Order, err error) {
	err = Db.Table("orders").Find(&orderlist).Error
	return orderlist, err
}