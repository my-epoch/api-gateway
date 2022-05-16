package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	service "github.com/my-epoch/api-gateway/gen/go/api/proto/v1"
	"github.com/my-epoch/api-gateway/pkg/consul"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"google.golang.org/grpc"
)

func registerServices(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	registerObjectService(mux, ctx, opts)
}

func registerObjectService(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	addr, err := consul.GetServiceAddr("object")
	if err != nil {
		logger.Warnf("cannot get object service addr: %e", err)
		return
	}

	if err = service.RegisterObjectServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		logger.Warnf("cannot register object service: %e", err)
		return
	}
}
