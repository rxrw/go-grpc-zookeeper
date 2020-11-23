package initializer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gitlab.dev.baai.ac.cn/basic-service/go-grpc-zookeeper/common"
	"gitlab.dev.baai.ac.cn/basic-service/go-grpc-zookeeper/registry"
	zk "gitlab.dev.baai.ac.cn/basic-service/go-grpc-zookeeper/registry/zookeeper"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//StartService 这里面registry负责将服务信息注册到zk，NewRpcServer启动本机服务
// 因此可能nodeId是个唯一键
func StartService(serviceName string, serviceVersion string, zkAddress []string, serverAddress string, serverPort int) (*grpc.Server, int) {

	instanceID := uuid.New()

	service := &registry.ServiceInfo{
		InstanceId: instanceID.String(),
		Name:       serviceName,
		Version:    serviceVersion,
		Address:    fmt.Sprintf("%s:%d", serverAddress, serverPort), //这个大概是服务端的地址吧。。就docker暴露出来的端口到宿主机，直接访问宿主机ip就ok
		Metadata:   metadata.Pairs(common.WeightKey, "1"),
	}

	registrar, err := zk.NewRegistrar(
		&zk.Config{
			ZkServers:      zkAddress, //这个是zk的地址们
			RegistryDir:    "/services",
			SessionTimeout: time.Second,
		})

	if err != nil {
		log.Panic(err)
		panic(err)
	}

	rs := grpc.NewServer()

	wg := sync.WaitGroup{}

	//run server的事情交给rd
	// wg.Add(1)
	// go func() {
	// 	Run(fmt.Sprintf("0.0.0.0:%d", serverPort), rs)
	// 	wg.Done()
	// }()

	wg.Add(1)
	go func() {
		registrar.Register(service)
		wg.Done()
	}()

	//这个进程监听关闭事件
	wg.Add(1)
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		registrar.Unregister(service)
		rs.GracefulStop()
		wg.Wait()
	}()

	return rs, serverPort
}
