package examples

import (
	"context"

	"gitlab.dev.baai.ac.cn/basic-service/go-grpc-zookeeper/proto"
)

//HelloWorldService make api a instance
type HelloWorldService struct {
	*proto.UnimplementedHelloWorldServer
}

//AreYouOk 我写了多少个注释了都
func (s *HelloWorldService) AreYouOk(ctx context.Context, in *proto.GreeterInfo) (*proto.ReturnMessage, error) {
	return &proto.ReturnMessage{
		Message: "I'm so fan",
	}, nil
}

//NewHelloWorldService 呵呵呵
func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}
