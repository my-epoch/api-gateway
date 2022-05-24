package service_config

import (
	"github.com/my-epoch/api-gateway/pkg/config"
)

type ServiceConfig struct {
	Name  string      `yaml:"name"`
	Addr  string      `yaml:"addr"`
	Check CheckConfig `yaml:"check"`
}

type CheckConfig struct {
	Interval string `yaml:"interval"`
	Timeout  string `yaml:"timeout"`
	GRPC     string `yaml:"GRPC"`
	HTTP     string `yaml:"HTTP"`
}

var serviceConfig *ServiceConfig

func Get() *ServiceConfig {
	if serviceConfig == nil {
		if err := config.LoadFile("config", &serviceConfig); err != nil {
			return nil
		}
	}

	return serviceConfig
}
