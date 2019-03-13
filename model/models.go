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