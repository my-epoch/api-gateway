package internal_config

import (
	"github.com/my-epoch/api-gateway/pkg/config"
	"github.com/my-epoch/api-gateway/pkg/logger"
)

type GatewayConfig struct {
	ListenAddr string `yaml:"listen_addr"`
}

var gwConfig *GatewayConfig

func GetGatewayConfig() *GatewayConfig {
	if gwConfig == nil {
		if err := config.LoadFile("config", &gwConfig); err != nil {
			logger.Warn("cannot load gateway config")
		}
	}
	return gwConfig
}
