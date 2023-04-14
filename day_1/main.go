package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"username": "name1", "data": "data1"})
	})
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run(":9090")

}
