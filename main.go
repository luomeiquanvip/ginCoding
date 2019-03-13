package main

import (
	"ginCoding/db"
	"ginCoding/router"
)

func main()  {
	db.Init()

	router := router.InitRouter()
	router.Run(":8888")

}
