package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r     := gin.Default()
	//后端接口
	admin := r.Group("/admin")
	{
		admin.POST("/goods/add",add)
		admin.POST("/goods/delete",delete)
		admin.POST("/goods/update",update)
		admin.POST("/goods/detail",detail)
		admin.POST("/goods/list",list)
		admin.GET("/goods/test",test)
	}
	//前端接口
	api := r.Group("/api")
	{
		api.POST("/goods/info",info)
	}
	r.Run(":9097")
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
//测试商品传递
func test(c * gin.Context){
	bn := c.DefaultQuery("title", "100")
	c.String(http.StatusOK, fmt.Sprintf("id= %s", bn))
}
func info(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"info",
		"msg":"info succ",
	})
}