// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldClient interface {
	AreYouOk(ctx context.Context, in *GreeterInfo, opts ...grpc.CallOption) (*ReturnMessage, error)
}

type helloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldClient(cc grpc.ClientConnInterface) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) AreYouOk(ctx context.Context, in *GreeterInfo, opts ...grpc.CallOption) (*ReturnMessage, error) {
	out := new(ReturnMessage)
	err := c.cc.Invoke(ctx, "/proto.HelloWorld/AreYouOk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
// All implementations must embed UnimplementedHelloWorldServer
// for forward compatibility
type HelloWorldServer interface {
	AreYouOk(context.Context, *GreeterInfo) (*ReturnMessage, error)
	mustEmbedUnimplementedHelloWorldServer()
}

// UnimplementedHelloWorldServer must be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (UnimplementedHelloWorldServer) AreYouOk(context.Context, *GreeterInfo) (*ReturnMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AreYouOk not implemented")
}
func (UnimplementedHelloWorldServer) mustEmbedUnimplementedHelloWorldServer() {}

// UnsafeHelloWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldServer will
// result in compilation errors.
type UnsafeHelloWorldServer interface {
	mustEmbedUnimplementedHelloWorldServer()
}

func RegisterHelloWorldServer(s grpc.ServiceRegistrar, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_AreYouOk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreeterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).AreYouOk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HelloWorld/AreYouOk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).AreYouOk(ctx, req.(*GreeterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AreYouOk",
			Handler:    _HelloWorld_AreYouOk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

// GoodbyeClient is the client API for Goodbye service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodbyeClient interface {
	SeeYouNever(ctx context.Context, in *GreeterInfo, opts ...grpc.CallOption) (*ReturnMessage, error)
}

type goodbyeClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodbyeClient(cc grpc.ClientConnInterface) GoodbyeClient {
	return &goodbyeClient{cc}
}

func (c *goodbyeClient) SeeYouNever(ctx context.Context, in *GreeterInfo, opts ...grpc.CallOption) (*ReturnMessage, error) {
	out := new(ReturnMessage)
	err := c.cc.Invoke(ctx, "/proto.Goodbye/SeeYouNever", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodbyeServer is the server API for Goodbye service.
// All implementations must embed UnimplementedGoodbyeServer
// for forward compatibility
type GoodbyeServer interface {
	SeeYouNever(context.Context, *GreeterInfo) (*ReturnMessage, error)
	mustEmbedUnimplementedGoodbyeServer()
}

// UnimplementedGoodbyeServer must be embedded to have forward compatible implementations.
type UnimplementedGoodbyeServer struct {
}

func (UnimplementedGoodbyeServer) SeeYouNever(context.Context, *GreeterInfo) (*ReturnMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeeYouNever not implemented")
}
func (UnimplementedGoodbyeServer) mustEmbedUnimplementedGoodbyeServer() {}

// UnsafeGoodbyeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodbyeServer will
// result in compilation errors.
type UnsafeGoodbyeServer interface {
	mustEmbedUnimplementedGoodbyeServer()
}

func RegisterGoodbyeServer(s grpc.ServiceRegistrar, srv GoodbyeServer) {
	s.RegisterService(&_Goodbye_serviceDesc, srv)
}

func _Goodbye_SeeYouNever_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreeterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodbyeServer).SeeYouNever(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Goodbye/SeeYouNever",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodbyeServer).SeeYouNever(ctx, req.(*GreeterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goodbye_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Goodbye",
	HandlerType: (*GoodbyeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeeYouNever",
			Handler:    _Goodbye_SeeYouNever_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
