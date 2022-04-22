package discovery

import (
	"github.com/FFFcomewhere/University-course-selection/pkg/common"
	"github.com/go-kit/kit/sd/etcdv3"
	"log"
	"sync"
)

type DiscoveryClientInstance struct {
	Host string //  Host
	Port string //  Port
	// 连接 consul 的配置
	config *etcdv3.ClientOptions
	client *etcdv3.Client
	mutex  sync.Mutex
	// 服务实例缓存字段
	instancesMap sync.Map
}

type DiscoveryClient interface {
	/**
	 * 服务注册接口
	 * @param serviceName 服务名
	 * @param instanceId 服务实例Id
	 * @param instancePort 服务实例端口
	 * @param weight 权重
	 * @param meta 服务实例元数据
	 */
	Register(serviceHost, servicePort, svcName, instanceId string, weight int, logger *log.Logger)

	/**
	 * 服务注销接口
	 * @param instanceId 服务实例Id
	 */
	DeRegister(instanceId string, logger *log.Logger)

	/**
	 * 发现服务实例接口
	 * @param serviceName 服务名
	 */
	DiscoveryServices(serviceName string, logger *log.Logger) []*common.ServiceInstance
}
