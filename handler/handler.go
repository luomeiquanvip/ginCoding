package handler

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}