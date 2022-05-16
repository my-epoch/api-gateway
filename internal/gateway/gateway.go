package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/my-epoch/api-gateway/internal/internal_config"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

var gwMux *runtime.ServeMux

func Init() {
	ctx := context.Background()

	gwMux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	registerServices(gwMux, ctx, opts)
}

func Serve() {
	gwConfig := internal_config.GetGatewayConfig()
	mux := http.NewServeMux()

	logger.Info("gateway starting")
	mux.Handle("/", gwMux)
	if err := http.ListenAndServe(gwConfig.ListenAddr, mux); err != nil {
		logger.Fatal("cannot start a server: %e", err)
	}
}
