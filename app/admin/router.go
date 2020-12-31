package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func checkToken(c * gin.Context){
	token := c.PostForm("token")
	fmt.Printf("token is    %v\n",token)
	c.Next()
}
func Routers(e * gin.Engine){
	e.Use(checkToken)
	goods := e.Group("/admin/goods")
	{
		goods.POST("/add",add)
		goods.POST("/delete",delete)
		goods.POST("/update",update)
		goods.POST("/detail",detail)
		goods.POST("/goods/list",list)
		goods.GET("/goods/test",test)
	}
	goodsType := e.Group("/admin/goods/type")
	{
		goodsType.POST("/add",typeAdd)
	}
}
func add(c * gin.Context){
	//bn := c.DefaultPostForm("bn","100")
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"add",
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
func typeAdd(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":"type add",
		"msg":"type add succ",
	})
}
