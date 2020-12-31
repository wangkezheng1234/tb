package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func Routers(e *gin.Engine){
	e.POST("/api/goods/info",info)
}
func info(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"info",
		"msg":"info succ",
	})
}