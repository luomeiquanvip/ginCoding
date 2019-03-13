package model

type Order struct {
	ID uint `json:"id"`
	User_name string `json:"user_name"`
	Amount string `json:"amount"`
	Status string `json:"status"`
	File_url int "json:file_url"
}
