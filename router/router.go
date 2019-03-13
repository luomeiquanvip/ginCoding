package router

import (
	."ginCoding/handler"
	"gopkg.in/gin-gonic/gin.v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//IndexApi为一个Handler
	router.GET("/", IndexApi)
	router.GET("/create", CreateAndInsert)
	router.GET("/orders", GetOrdersApi)
	router.GET("/update", UpdateOrder)
	router.GET("/fuzzySearch", FuzzySearch)


	return router
}
