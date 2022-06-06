package main

import (
	"github.com/my-epoch/api-gateway/internal/gateway"
	"github.com/my-epoch/api-gateway/pkg/consul_client"
	"github.com/my-epoch/api-gateway/pkg/logger"
)

func main() {
	logger.Info("initializing")
	consul_client.InitClient()

	gateway.Init()
	gateway.Serve()
}
