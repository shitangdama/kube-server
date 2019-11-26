package v1

import (
	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
