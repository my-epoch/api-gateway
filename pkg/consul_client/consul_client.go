package consul_client

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/my-epoch/api-gateway/pkg/logger"
)

var client *consulApi.Client

func InitClient() {
	consulConfig := consulApi.DefaultConfig()
	consulConfig.Address = "consul.service.consul:8500"
	var err error
	client, err = consulApi.NewClient(consulConfig)
	if err != nil {
		logger.Fatal("cannot create Consul client: %e", err)
	}
}
