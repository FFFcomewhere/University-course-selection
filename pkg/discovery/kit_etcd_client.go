package discovery

import (
	"context"
	"fmt"
	"github.com/FFFcomewhere/University-course-selection/pkg/common"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"strconv"
	"strings"
	"time"
)

//创建etcd客户端
func MakeEtcdClient(etcdHost string, etcdPort string) *DiscoveryClientInstance {
	clientOptions := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}
	machines := []string{etcdHost + ":" + etcdPort}
	client, err := etcdv3.NewClient(context.Background(), machines, clientOptions)

	if err != nil {
		fmt.Println("etcd connect err: ", err)
	}

	return &DiscoveryClientInstance{
		Host:   etcdHost,
		Port:   etcdPort,
		config: &clientOptions,
		client: &client,
	}
}

func (etcdClient *DiscoveryClientInstance) Register(serviceHost, servicePort, svcName, instanceId string, weight int, logger *log.Logger) {

	key := instanceId
	value := serviceHost + ":" + servicePort + "/" + svcName + "/"

	// 创建注册器
	register := etcdv3.NewRegistrar(*etcdClient.client, etcdv3.Service{
		Key:   key,
		Value: value,
	}, log.NewNopLogger())

	// 注册器启动注册
	register.Register()
}

func (etcdClient *DiscoveryClientInstance) DeRegister(instanceId string, logger *log.Logger) {
	// 查询服务对于的主机号和端口号
	client := *etcdClient.client
	key := instanceId
	values, err := client.GetEntries(key)

	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}

	fmt.Println(values)
	//// 创建注册器
	//register := etcdv3.NewRegistrar(*etcdClient.client, etcdv3.Service{
	//	Key:   instanceId,
	//	Value: values[0],
	//}, log.NewNopLogger())
	//
	////注销服务
	//register.Deregister()
}

//TODO 还未完成
func (etcdClient *DiscoveryClientInstance) DiscoveryServices(serviceName string, logger *log.Logger) []*common.ServiceInstance {
	//  该服务已监控并缓存
	instanceList, ok := etcdClient.instancesMap.Load(serviceName)
	if ok {
		return instanceList.([]*common.ServiceInstance)
	}

	// 申请锁
	etcdClient.mutex.Lock()
	defer etcdClient.mutex.Unlock()
	// 再次检查是否监控
	instanceList, ok = etcdClient.instancesMap.Load(serviceName)
	if ok {
		return instanceList.([]*common.ServiceInstance)
	} else {
		// 注册监控
		go func() {
			//创建实例管理器, 此管理器会Watch监听etc中prefix的目录变化更新缓存的服务实例数据

			//logger := log.NewNopLogger()
			//instancer, err := etcdv3.NewInstancer(*etcdClient.client, serviceName, logger)
			//if err != nil {
			//	fmt.Println("开启实例管理器失败: ", err)
			//}

		}()

	}

	return []*common.ServiceInstance{}
}

func newServiceInstance(service *etcdv3.Service, logger *log.Logger) *common.ServiceInstance {
	index1 := strings.Index(service.Value, ":")
	index2 := strings.Index(service.Value, "/")

	host := service.Value[0 : index1+1]
	port := strconv.Atoi(service.Value[index1 : index2+1])
	rpcPort := port - 1

	return &common.ServiceInstance{
		Host:     host,
		Port:     port,
		GrpcPort: rpcPort,
		Weight:   nil, //TODO 暂时没有使用带权重的平滑轮询策略
	}
}
