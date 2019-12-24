/*
/*
@Time : 2019/12/23 2:56 下午
@Author : tianpeng.du
@File : main
@Software: GoLand
*/
package main

import (
	"github.com/du2016/code-generator/pkg/client/clientset/versioned"
	"github.com/du2016/code-generator/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/tianpeng.du/kubeconfig")
	if err != nil {
		log.Println(err)
		return
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		log.Println(err)
		return
	}
	informer := externalversions.NewSharedInformerFactoryWithOptions(clientset, 10*time.Second, externalversions.WithNamespace("default"))
	go informer.Start(nil)
	//time.Sleep(10*time.Second)

	IpCrdInformer:=informer.Rocdu().V1().Ips()
	cache.WaitForCacheSync(nil,IpCrdInformer.Informer().HasSynced)

	for {
		ips, err := IpCrdInformer.Lister().List(labels.Everything())
		if err != nil {
			log.Println(err)
		}
		for _,v:=range ips{
			log.Println(v.Name)
		}
		time.Sleep(1*time.Second)
	}
}
