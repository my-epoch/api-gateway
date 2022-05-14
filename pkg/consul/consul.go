package consul

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/api-gateway/pkg/service_config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var client *consulApi.Client
var cfg *service_config.ServiceConfig

func RegisterService(serviceCfg *service_config.ServiceConfig) {
	if serviceCfg == nil {
		logger.Fatal("cannot register service: ServiceConfig is nil")
	}
	cfg = serviceCfg
	var err error
	client, err = consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		logger.Fatal("cannot create Consul client: %e", err)
	}

	if err := client.Agent().ServiceRegister(
		&consulApi.AgentServiceRegistration{
			Name:    cfg.Name,
			Address: cfg.Addr,
			Check: &consulApi.AgentServiceCheck{
				GRPC:     cfg.Addr,
				Interval: cfg.Check.Interval,
				Timeout:  cfg.Check.Timeout,
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

func GetServiceAddr(name string) (string, error) {
	addrs, _, err := client.Health().Service(name, "", true, nil)
	if err != nil {
		return "", err
	}

	if len(addrs) < 1 {
		return "", status.Error(codes.NotFound, "service not found")
	}

	return addrs[0].Service.Address, nil
}

func DeregisterService() {
	if client == nil {
		logger.Warn("cannot deregister service: service is not registered")
	}
	if err := client.Agent().ServiceDeregister(cfg.Name); err != nil {
		logger.Fatalf("cannot deregister service: %e", err)
	}
}
