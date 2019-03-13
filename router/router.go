package router

import (
	."ginCoding/handler"
	"gopkg.in/gin-gonic/gin.v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//IndexApi为一个Handler
	router.GET("/", IndexApi)
	router.GET("/createAndInsert", CreateAndInsert)
	router.GET("/Orders", Orders)
	router.GET("/updateOrder", UpdateOrder)
	router.GET("/fuzzySearch", FuzzySearch)
	router.GET("/infoOrderByAmount", InfoOrderByAmount)
	router.GET("/txInsert", TxInsert)
	router.POST("/upload", UploadFile)


	return router
}
