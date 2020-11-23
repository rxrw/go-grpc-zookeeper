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
Then set `export CONFIG_FILE=/path/to/your/config.yml`
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

# Usage
### Configuration
A sample configuration is like this:
```yml
name: test # Your Service Name
port: 5001 # Your Service Port
address: 127.0.0.1 # Your Host Ip
version: 1.0 # Your Service Version

discovery:
  url: # zookeeper addresses
    - 127.0.0.1
  port: 2181 # zookeeper ports
  type: zookeeper 

client:
  balancer: random #consistent_hash, ketama, least_connection, random, round_robin
  servers:
    - name: "test" # as a client, this is your  dependencies service name
      version: "1.0" # version
  insecure: false # no use

```

### Server And Client
For developing an app with both server and client available, we should have these config completed and we have an instance of `*grpc.Server` returned. You should register your service with its **proto register function** like `proto.RegisterHelloWorldService(*grpc.Server, &HelloWorldService)`.
In some case like using iris, it can support original gRPC either. So we can use grpc like [this](https://github.com/kataras/iris/wiki/Grpc)

# License
MIT
