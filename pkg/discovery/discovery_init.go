package discovery

import (
	"errors"
	"github.com/FFFcomewhere/University-course-selection/pkg/bootstrap"
	"github.com/FFFcomewhere/University-course-selection/pkg/common"
	"github.com/FFFcomewhere/University-course-selection/pkg/loadbalance"
	uuid "github.com/satori/go.uuid"

	"log"
	"os"
)

var EtcdService DiscoveryClient
var LoadBalance loadbalance.LoadBalance
var Logger *log.Logger
var NoInstanceExistedErr = errors.New("no available client")

func init() {
	// 1.实例化一个 Etcd 客户端，
	EtcdService = MakeEtcdClient(bootstrap.DiscoverConfig.Host, bootstrap.DiscoverConfig.Port)
	LoadBalance = new(loadbalance.RandomLoadBalance)
	Logger = log.New(os.Stderr, "", log.LstdFlags)
}

func Register() {
	//// 实例失败，停止服务
	if EtcdService == nil {
		panic(0)
	}

	//判空 instanceId,通过 go.uuid 获取一个服务实例ID
	instanceId := bootstrap.DiscoverConfig.InstanceId

	if instanceId == "" {
		instanceId = bootstrap.DiscoverConfig.ServiceName + uuid.NewV4().String()
	}

	//注册服务
	EtcdService.Register(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port, instanceId, instanceId,
		bootstrap.DiscoverConfig.Weight, Logger)
	//TODO 这里最好再添加一个panic

	Logger.Printf(bootstrap.DiscoverConfig.ServiceName+"-service for service %s success.", bootstrap.DiscoverConfig.ServiceName)

}

func Deregister() {
	//// 实例失败，停止服务
	if EtcdService == nil {
		panic(0)
	}

	//判空 instanceId,通过 go.uuid 获取一个服务实例ID
	instanceId := bootstrap.DiscoverConfig.InstanceId

	if instanceId == "" {
		instanceId = bootstrap.DiscoverConfig.ServiceName + "-" + uuid.NewV4().String()
	}

	EtcdService.DeRegister(instanceId, Logger)
}

func DiscoveryService(serviceName string) (*common.ServiceInstance, error) {
	instances := EtcdService.DiscoveryServices(serviceName, Logger)

	if len(instances) < 1 {
		Logger.Printf("no available client for %s.", serviceName)
		return nil, NoInstanceExistedErr
	}
	return LoadBalance.SelectService(instances)
}
