package v1

import (
	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetRepo(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "test",
	})
}

// //获取单个文章
// func GetArticle(c *gin.Context) {
// }

// //获取多个文章
// func GetArticles(c *gin.Context) {
// }

// //新增文章
// func AddArticle(c *gin.Context) {
// }

// //修改文章
// func EditArticle(c *gin.Context) {
// }

// //删除文章
// func DeleteArticle(c *gin.Context) {
// }
