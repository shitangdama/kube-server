package v1

import (
	"laki/common"

	"github.com/gin-gonic/gin"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	clientset *kubernetes.Clientset
	podsList  *core_v1.PodList
	err       error
)

//获取多个文章标签
func GetPod(c *gin.Context) {
	clientset, _ = common.InitClient()
	podsList, _ = clientset.CoreV1().Pods("default").List(meta_v1.ListOptions{})

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": podsList,
	})
}

func PostPod(c *gin.Context) {
	// clientset, _ = common.InitClient()
	// podsList, _ = clientset.CoreV1().Pods("default").List(meta_v1.ListOptions{})

	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "success",
		// "data": podsList,
	})
}
