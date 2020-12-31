package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e * gin.Engine){
	e.POST("/admin/goods/add",add)
	e.POST("/admin/goods/delete",delete)
	e.POST("/admin/goods/update",update)
	e.POST("/admin/goods/detail",detail)
	e.POST("/admin/goods/list",list)
	e.GET("/admin/goods/test",test)
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