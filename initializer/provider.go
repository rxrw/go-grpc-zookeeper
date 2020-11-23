package initializer

import (
	"fmt"
	"log"
	"net"

	"github.com/rxrw/go-grpc-zookeeper/config"

	"google.golang.org/grpc"
)

func registerAndStartServer() (*grpc.Server, int) {
	serverConfig := config.GetRPCConfig(false)

	return StartService(serverConfig.Name, serverConfig.Version, serverConfig.Discovery.Url, serverConfig.Address, serverConfig.Port)
}

func registerClient() map[string]*grpc.ClientConn {
	clientConfig := config.GetRPCConfig(false)
	Conn := make(map[string]*grpc.ClientConn)
	// 多例注册到数组
	for _, service := range clientConfig.Client.Servers {
		Conn[service.Name] = RegisterClient(clientConfig.Discovery.Url, service.Name, service.Version)
	}
	return Conn
}

//Provide 提供grpc服务，注册到zk
func Provide() (*grpc.Server, net.Listener) {
	registerClient()
	rs, port := registerAndStartServer()
	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))

	return rs, listener
}

//Run 运行这个server
func Run() {

	//! 记得serve前要注册服务的

	rs, listener := Provide()

	err := rs.Serve(listener)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

}
