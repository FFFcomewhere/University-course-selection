package main

import (
	"GolandProjects/base_study/etcd/demo1/config"
	"GolandProjects/base_study/etcd/demo1/discovery"
	"GolandProjects/base_study/etcd/demo1/endpoint"
	"GolandProjects/base_study/etcd/demo1/plugins"
	"GolandProjects/base_study/etcd/demo1/service"
	"GolandProjects/base_study/etcd/demo1/transport"
	"context"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	var (
		etcdHost    = flag.String("etcd_host", "127.0.0.1", "etcd host")
		etcdPort    = flag.Int("etcd_port", 2379, "etcd port")
		serviceHost = flag.String("service_host", "127.0.0.1", "service host")
		servicePort = flag.Int("service_port", 10086, "service port")
		serviceName = flag.String("service_name", "string", "service name")
	)
	flag.Parse()

	ctx := context.Background()
	var svc service.Service
	svc = service.StringService{}
	//add logging middleware
	svc = plugins.LoggingMiddleware(config.KitLogger)(svc)

	stringEndpoint := endpoint.MakeStringEndpoint(svc)
	healthEndpoint := endpoint.MakeHealthCheckEndpoint(svc)
	sayEndpoint := endpoint.MakeSayEndpoint(svc)

	endpts := endpoint.StringEndpoints{
		StringEndpoint:      stringEndpoint,
		HealthCheckEndpoint: healthEndpoint,
		SayEndpoint:         sayEndpoint,
	}

	//new http.Handler
	r := transport.MakeHttpHandler(ctx, endpts)

	//开启etcd服务器
	var registerClient discovery.RegisterClient
	registerClient, err := discovery.NewRegisterClient(*etcdHost, *etcdPort)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		config.Logger.Println("http server start at port: " + strconv.Itoa(*servicePort))
		//注册服务
		key := "/" + *serviceName
		value := *serviceHost + ":" + strconv.Itoa(*servicePort)
		fmt.Println(key, value)
		registerClient.Register(key, value, 5)

		//监听续租相应chan
		go registerClient.ListenLeasRespChan()
		select {}
	}()

	handler := r
	http.ListenAndServe(":"+strconv.Itoa(*servicePort), handler)
}
