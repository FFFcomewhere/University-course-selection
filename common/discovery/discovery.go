package discovery

import (
	"context"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"log"
	"sync"
	"time"
)

type ServiceDiscovery struct {
	client     *clientv3.Client
	serverList map[string]string //服务列表
	lock       sync.Mutex
}

//新建发现服务
func NewServiceDiscovery(endpoints []string) *ServiceDiscovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &ServiceDiscovery{
		client:     client,
		serverList: make(map[string]string),
	}
}

//初始化服务列表和监视
func (s *ServiceDiscovery) WatchService(prefix string) error {
	//根据前缀获取现有key
	resp, err := s.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}

	//监视前缀,修改变更的server
	go s.watcher(prefix)
	return nil
}

//watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
	rch := s.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching prefix:%s now... \n", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: //删除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

// 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = string(value)
	log.Println("put key:", key, "value:", value)
}

//删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Println("del key:", key)
}

//获取服务地址
func (s *ServiceDiscovery) GetSService() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, 0)

	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

//关闭服务
func (s *ServiceDiscovery) Close() error {
	return s.client.Close()
}

//
//func main() {
//	var endpoints = []string{"localhost:2379"}
//	service := NewServiceDiscovery(endpoints)
//	defer service.Close()
//	service.WatchService("/web/")
//	service.WatchService("/gRPC")
//	service.WatchService("/time")
//	for {
//		select {
//		case <-time.Tick(10 * time.Second):
//			log.Println(service.GetSService())
//		}
//	}
//}
