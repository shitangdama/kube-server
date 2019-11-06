package main

import (
	// "bufio"
	"flag"
	"fmt"
	"path/filepath"

	// "os"
	// apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	// "k8s.io/client-go/util/retry"
)

func main() {

	var kubeconfig *string
	fmt.Println(homedir.HomeDir())
	kubeconfig = flag.String("kubeconfig", filepath.Join("./", "kubeconfig.yml"), "(optional) absolute path to the kubeconfig file")
	fmt.Println(kubeconfig)

	namespace := "default"
	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	list, err := client.Resource(deploymentRes).Namespace(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(list.Items)
	for k := range list.Items {
		fmt.Println(k)
	}

}
