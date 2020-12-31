package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r := gin.Default()
	r.POST("/admin/goods/add",add)
	r.POST("/admin/goods/delete",delete)
	r.POST("/admin/goods/update",update)
	r.POST("/admin/goods/detail",detail)
	r.POST("/admin/goods/list",list)
	r.Run(":9095")
}
func add(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"add",
		"msg":"add succ",
	})
}
func delete(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"delete",
		"msg":"delete succ",
	})
}
func update(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"update",
		"msg":"update succ",
	})
}
func detail(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"detail",
		"msg":"detail succ",
	})
}
func list(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"list",
		"msg":"list succ",
	})
}