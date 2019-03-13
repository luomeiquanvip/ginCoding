package model

import (
	."ginCoding/db"
)

type Order struct {
	ID uint `json:"id"`
	Order_id  string `json:"order_id"`
	User_name string `json:"user_name"`
	Amount float64 `json:"amount"`
	Status string `json:"status"`
	File_url string  `json:"file_url"`
}

func (p *Order) GetOrders() (orders []Order, err error) {
	Db.Table("orders").Find(&orders)
	return
}

func (p *Order) GetFuzzySearchs() (orders []Order, err error) {
	Db.Where("order_id LIKE ?","%2%").Find(&orders)
	return
}

func (p *Order) GetInfoOrderByAmount() (orders []Order, err error) {
	Db.Order("amount desc").Find(&orders)
	return
}

func (p *Order) TxInsert() (orders []Order, err error) {
	//插入数据
	order1 := Order{Order_id: "444", User_name: "abc", Amount: 44.4,Status:"y",File_url:"www.qiyun.com"}
	order2 := Order{Order_id: "555", User_name: "def", Amount: 55.5,Status:"n",File_url:"www.shanghai.com"}
	order3 := Order{Order_id: "666", User_name: "ghi", Amount: 66.6,Status:"y",File_url:"www.shengzhen.com"}

	tx := Db.Begin()
	if err1 := tx.Create(&order1).Error; err1 != nil {
		tx.Rollback()
		err = err1
		return
	}
	orders = append(orders,order1)

	if err2 := tx.Create(&order2).Error; err2 != nil {
		tx.Rollback()
		err = err2
		return
	}
	orders = append(orders,order2)

	if err3 := tx.Create(&order3).Error; err3 != nil {
		tx.Rollback()
		err = err3
		return
	}
	orders = append(orders,order3)

	tx.Commit()
	return orders,nil

}