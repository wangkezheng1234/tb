package main

import (
	"fmt"
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
	r.GET("/admin/goods/test",test)
	r.Run(":9096")
}
func add(c * gin.Context){
	bn := c.DefaultPostForm("bn","100")
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":bn,
		"msg":"add goods succ",
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
func test(c * gin.Context){
	name := c.DefaultQuery("title", "枯藤")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}