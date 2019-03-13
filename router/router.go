package router

import (
	."ginCoding/handler"
	"gopkg.in/gin-gonic/gin.v1"
)
func InitRouter() *gin.Engine {
	router := gin.Default()
	//IndexApi为一个Handler
	router.GET("/", IndexApi)

	return router
}
