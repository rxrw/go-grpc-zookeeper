package initializer

import (
	"log"

	"gitlab.dev.baai.ac.cn/basic-service/go-grpc-zookeeper/balancer"

	registry "gitlab.dev.baai.ac.cn/basic-service/go-grpc-zookeeper/registry/zookeeper"

	"google.golang.org/grpc"
)

//RegisterClient 注册客户端
func RegisterClient(zkServer []string, serviceName string, serviceVersion string) *grpc.ClientConn {
	registry.RegisterResolver("zk", zkServer, "/services", serviceName, serviceVersion)
	c, err := grpc.Dial("zk:///", grpc.WithInsecure(), grpc.WithBalancerName(balancer.RoundRobin))
	if err != nil {
		log.Fatal(err)
	}

	return c
}
