package discovery

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"log"
	"strconv"
	"testing"
)

func TestRegister(t *testing.T) {

	var (
		etcdHost    = flag.String("etcd-host", "127.0.0.1", "etcd host")
		etcdPort    = flag.Int("etcd-port", 2379, "etcd port")
		serviceHost = flag.String("service-host", "127.0.0.1", "service host")
		servicePort = flag.Int("service-port", 10086, "service port")
		serviceName = flag.String("service-name", "string", "service name")
	)

	//开启etcd服务器
	var registerClient RegisterClient
	registerClient, err := NewRegisterClient(*etcdHost, *etcdPort)
	if err != nil {
		log.Println(err)
	}

	//注册服务
	key := "/" + *serviceName
	value := *serviceHost + ":" + strconv.Itoa(*servicePort)
	fmt.Println(key, value)
	registerClient.Register(key, value, 5)

	//监听续租相应chan
	go registerClient.ListenLeasRespChan()
	select {}

}

func TestEtcdPushData(t *testing.T) {
	var (
		etcdHost = flag.String("etcd-host", "127.0.0.1", "etcd host")
		etcdPort = flag.Int("etcd-port", 2379, "etcd port")
	)

	//开启etcd服务器
	//var registerClient RegisterClient
	//registerClient, err := NewRegisterClient(*etcdHost, *etcdPort)
	//if err != nil {
	//	fmt.Println(err)
	//}
	var etcdAddrs = []string{*etcdHost + ":" + strconv.Itoa(*etcdPort)}
	var options = etcdv3.ClientOptions{DialTimeout: ttl}
	client := etcdv3.NewClient(context.Background(), etcdAddrs)

}
