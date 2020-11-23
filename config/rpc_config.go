package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

//RPCConfig rpc的配置.
type RPCConfig struct {
	Name      string
	Port      int
	Address   string
	Version   string
	Discovery discoveryConfig
	Client    clientConfig
}

//DiscoveryConfig 服务发现的配置
type discoveryConfig struct {
	Url  []string
	Port int
	Type string
}

//ClientConfig 作为客户端的配置
type clientConfig struct {
	Balancer string
	Servers  []serverItem
	Insecure bool
}

//ServiceItem 每个服务节点配置
type serverItem struct {
	Name    string
	Version string
}

var mu sync.Mutex

var config *RPCConfig

//GetRPCConfig 获取rpc的配置
func GetRPCConfig(update bool) *RPCConfig {
	if config != nil && !update {
		return config
	}

	mu.Lock()
	defer mu.Unlock()

	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = "/app/rpc-config.yml"
	}

	data, _ := ioutil.ReadFile(configFile)

	t := RPCConfig{}

	yaml.Unmarshal(data, &t)

	fmt.Println("Reading Confg: ", t)

	config = &t

	return config
}
