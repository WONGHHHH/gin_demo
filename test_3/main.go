package main

import (

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {

	NestedStruct StructA

	FieldB string `form:"field_b"`
}

type StructC struct {

	NestStructPionter *StructA

	FieldC string `form:"field_c"`
}

type StructD struct {

	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}

	FieldD string `form:"field_d"`
}

func GetDateB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{

		"a" : b.FieldB,
		"b" : b.NestedStruct,
	})
}

func GetDateC(c *gin.Context) {
	var sc StructC
	c.Bind(&sc)
	c.JSON(200, gin.H{

		"a" : sc.FieldC,
		"b" : sc.NestStructPionter,
	})
}

func GetDateD(c *gin.Context) {
	var sd StructD
	c.Bind(&sd)
	c.JSON(200, gin.H{

		"a" : sd.FieldD,
		"b" : sd.NestedAnonyStruct,
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	r.GET("/getB", GetDateB)
	r.GET("/getC", GetDateC)
	r.GET("/getD", GetDateD)
	
	
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run(":9090")
	// curl "http://localhost:9090/getB?field_a=hello&field_b=word"
	// curl "http://localhost:9090/getC?field_a=hello&field_c=word"
	// curl "http://localhost:9090/getD?field_x=hello&field_d=word"
}
