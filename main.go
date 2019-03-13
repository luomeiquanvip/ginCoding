package main

import (
	db "ginCoding/db"
	. "ginCoding/router"
)

func main()  {
	defer db.Db.Close()
	db.Init()

	router := InitRouter()
	router.Run(":8888")

}
