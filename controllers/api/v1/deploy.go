package v1

import (
	"fmt"
	"laki/common"

	"github.com/gin-gonic/gin"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//获取多个文章标签
func GetDeploys(c *gin.Context) {
	clientset, _ := common.InitClient()
	deploymentList, _ := clientset.AppsV1().Deployments("default").List(meta_v1.ListOptions{})
	// .Deployments("default").List(meta_v1.ListOptions{})

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": deploymentList,
	})
}

//获取1ge
func GetDeploy(c *gin.Context) {

	ns := "default"
	if c.Query("ns") != "" {
		ns = c.Query("ns")
	}

	name := c.Param("name")
	fmt.Println(name)

	clientset, _ := common.InitClient()
	if deploy, err := clientset.AppsV1().Deployments(ns).Get(name, meta_v1.GetOptions{}); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "fail",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": deploy,
		})
	}

}
