package discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"strconv"
	"time"
)

type RegisterClient struct {
	Host    string // etcd  host
	Port    int    //etcd port
	client  *clientv3.Client
	lesseID clientv3.LeaseID // 租约id
	//租约keepalieve相关chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
	value         string
}

func NewRegisterClient(etcdHost string, etcdPort int) (RegisterClient, error) {
	var endpoints = []string{etcdHost + ":" + strconv.Itoa(etcdPort)}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return RegisterClient{}, err
	}

	return RegisterClient{
		Host:   etcdHost,
		Port:   etcdPort,
		client: client,
	}, err

}

func (registerClient *RegisterClient) Register(key, value string, lease int64) error {
	//在Key后面添加序列号,防止Key值被覆盖
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := registerClient.client.Get(ctx, key, clientv3.WithPrefix())
	cancel()

	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return err
	}

	key = key + strconv.Itoa(int(resp.Count))

	registerClient.key = key
	registerClient.value = value

	//申请租约设置时间keeplive
	if err := registerClient.putKeyWithLease(lease); err != nil {
		return err
	}

	return nil
}

//注销服务
func (registerClient *RegisterClient) DeRegister() error {
	//撤销租约
	if _, err := registerClient.client.Revoke(context.Background(), registerClient.lesseID); err != nil {
		return err
	}

	log.Panicln("撤销租约")
	return registerClient.client.Close()
}

//设置租约
func (registerClient *RegisterClient) putKeyWithLease(lease int64) error {
	//设置租约时间
	resp, err := registerClient.client.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	//注册服务并绑定租约
	_, err = registerClient.client.Put(context.Background(), registerClient.key, registerClient.value, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}

	//设置续租 定期发送需求请求
	leaseRespChan, err := registerClient.client.KeepAlive(context.Background(), resp.ID)

	if err != nil {
		return err
	}
	registerClient.lesseID = resp.ID
	log.Println(registerClient.lesseID)
	registerClient.keepAliveChan = leaseRespChan
	log.Printf("Put key:%s value:%s success! \n", registerClient.key, registerClient.value)
	return nil
}

//监听 续租情况
func (registerClient *RegisterClient) ListenLeasRespChan() {
	for leaseKeepResp := range registerClient.keepAliveChan {
		log.Println("续约成功", leaseKeepResp)
	}

	log.Println("关闭续租")
}

//func main() {
//
//	var (
//		etcdHost    = flag.String("etcd-host", "127.0.0.1", "etcd host")
//		etcdPort    = flag.Int("etcd-port", 2379, "etcd port")
//		serviceHost = flag.String("service-host", "127.0.0.1", "service host")
//		servicePort = flag.Int("service-port", 10086, "service port")
//		serviceName = flag.String("service-name", "string", "service name")
//	)
//
//	//开启etcd服务器
//	var registerClient RegisterClient
//	registerClient, err := NewRegisterClient(*etcdHost, *etcdPort)
//	if err != nil {
//		log.Println(err)
//	}
//
//	//注册服务
//	key := "/" + *serviceName
//	value := *serviceHost + ":" + strconv.Itoa(*servicePort)
//	fmt.Println(key, value)
//	registerClient.Register(key, value, 5)
//
//	//监听续租相应chan
//	go registerClient.ListenLeasRespChan()
//	select {}
//}
