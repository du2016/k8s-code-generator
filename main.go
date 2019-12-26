/*
/*
@Time : 2019/12/23 2:56 下午
@Author : tianpeng.du
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	ip_v1 "github.com/du2016/code-generator/pkg/apis/ip/v1"
	"github.com/du2016/code-generator/pkg/client/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"log"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/tianpeng.du/.kube/config")
	if err != nil {
		log.Println(err)
		return
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		log.Println(err)
		return
	}
	//informer := externalversions.NewSharedInformerFactoryWithOptions(clientset, 10*time.Second, externalversions.WithNamespace("default"))
	//go informer.Start(nil)
	////time.Sleep(10*time.Second)
	//
	//IpCrdInformer:=informer.Ip().V1().Ips()
	//cache.WaitForCacheSync(nil,IpCrdInformer.Informer().HasSynced)

	//for {
	//	ips, err := IpCrdInformer.Lister().List(labels.Everything())
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	for _,v:=range ips{
	//		log.Println(v.Name)
	//	}
	//	time.Sleep(1*time.Second)
	//}
	ipc := IPcontroller{}
	watchList := cache.NewListWatchFromClient(clientset.IpV1().RESTClient(), "ips", v1.NamespaceAll, fields.Everything())
	ipStore, ipController := cache.NewInformer(
		watchList,
		&v1.Pod{},
		time.Second*30,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    ipc.enqueueIp,
			UpdateFunc: ipc.updatePod,
		},
	)
	ipc = IPcontroller{
		Store:      ipStore,
		Controller: ipController,
		Queue:      workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ips"),
	}
	cache.WaitForCacheSync(nil, ipc.Controller.HasSynced)

	var stopCh <-chan struct{}
	for i := 0; i < 10; i++ {
		go wait.Until(ipc.worker, time.Second, stopCh)
	}

	<-stopCh

	ipc.Queue.ShutDown()
}

type IPcontroller struct {
	Store      cache.Store
	Controller cache.Controller
	Queue      workqueue.RateLimitingInterface
}

func (self *IPcontroller) enqueueIp(obj interface{}) {
	log.Println(obj.(ip_v1.Ip))
	self.Queue.Add(obj)
}

func (self *IPcontroller) updatePod(obj interface{}, new interface{}) {
	log.Println(obj.(ip_v1.Ip), new)
	self.enqueueIp(new)
}

func (self *IPcontroller) worker() {
	workFunc := func() bool {
		key, quit := self.Queue.Get()
		if quit {
			return true
		}
		self.Store.Resync()
		obj, exists, err := self.Store.GetByKey(key.(string))
		if !exists {
			fmt.Printf("Pod has been deleted %v\n", key)
			return false
		}
		if err != nil {
			fmt.Printf("cannot get pod: %v\n", key)
			return false
		}
		log.Println(obj)
		return false
	}

	for {
		if quit := workFunc(); quit {
			fmt.Printf("pod worker shutting down")
			return
		}
	}
}
