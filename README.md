# go-grpc-zookeeper
A Client that Connect With Zookeeper for gRPC.

## Inspired
[liyue201/grpc-lb](https://github.com/etcd-manage/etcd-manage-server): On the base of this project.

## Introduction
An library for go and grpc to use zookeeper as its services discovery and balancer.
The origin project has the support for etcd and consul, but as we know, the version control of etcd is so massy that we have to drop it.
This project can make your service registered on zookeeper and call it with just service name. Just like dubbo with zookeeper in java.

## Quick Start
1. You should have a zookeeper instance running.
```bash
docker run -p 2181:2181 --privileged --name zk zookeeper:latest
```
2. Put `rpc-config.yml.example` into any place with permissions we can read and rename to `rpc-config.yml`
You can modify your config file as what you like.
3. create a new project with go modules.
```bash
mkdir demo
cd demo
go mod init demo
```
4. create a file with the name `main.go` and run command:
```bash
go get -u github.com/rxrw/go-grpc-zookeeper
```
5. make your `main.go` like this:
```go
func main() {
    app,listener := initializer.Provide()
    proto.RegisterHelloWorldServer(app, &examples.HelloWorldService)
    app.Run()
}
```
6. Just 
```bash
go run main.go
```
7. Let's look at the services provided:
```bash
docker exec -it zk ./bin/zkCli.sh
cd services
ls /services/test/1.0
# Now You can See An Uuid listed
# ls it you can see your ip and port registered on zookeeper.
```
8. Stop the `main.go` and the key above will disappear

> You can also start an other project and change the yml `name` to run the client. It's the same as normal grpc.

# Licence
MIT
