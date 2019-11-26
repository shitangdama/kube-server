package controllers

import (
	v1 "laki/controllers/api/v1"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiv1 := r.Group("/api/v1")

	{
		//获取标签列表
		apiv1.GET("/info", v1.GetInfo)
		apiv1.GET("/pods", v1.GetPod)
	}
	return r
}
