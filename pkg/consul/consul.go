package consul

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/api-gateway/pkg/service_config"
)

var client *consulApi.Client
var cfg *service_config.ServiceConfig

func InitClient() {
	consulConfig := consulApi.DefaultConfig()
	consulConfig.Address = "consul.service.consul:8500"
	var err error
	client, err = consulApi.NewClient(consulConfig)
	if err != nil {
		logger.Fatal("cannot create Consul client: %e", err)
	}
}

func RegisterService(serviceCfg *service_config.ServiceConfig) {
	if serviceCfg == nil {
		logger.Fatal("cannot register service: ServiceConfig is nil")
	}
	cfg = serviceCfg

	InitClient()
	if err := client.Agent().ServiceRegister(
		&consulApi.AgentServiceRegistration{
			Name:    cfg.Name,
			Address: cfg.Addr,
			Check: &consulApi.AgentServiceCheck{
				GRPC:     cfg.Check.GRPC,
				Interval: cfg.Check.Interval,
				Timeout:  cfg.Check.Timeout,
				HTTP:     cfg.Check.HTTP,
			},
			Connect: &consulApi.AgentServiceConnect{
				Native: true,
			},
		},
	); err != nil {
		logger.Fatal("cannot register service: %e", err)
	}
	logger.Info("service successful registered in Consul")
}

func DeregisterService() {
	if client == nil {
		logger.Warn("cannot deregister service: service is not registered")
	}
	if err := client.Agent().ServiceDeregister(cfg.Name); err != nil {
		logger.Fatalf("cannot deregister service: %e", err)
	}
}
