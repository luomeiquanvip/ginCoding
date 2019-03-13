package model

import (
	db"ginCoding/db"
)

type Order struct {
	ID uint `json:"id"`
	Order_id  string `json:"order_id"`
	User_name string `json:"user_name"`
	Amount float64 `json:"amount"`
	Status string `json:"status"`
	File_url string  `json:"file_url"`
}

func (p *Order) GetPersons() (persons []Order, err error) {
	db.Db.Table("orders").Find(&persons)
	return
}