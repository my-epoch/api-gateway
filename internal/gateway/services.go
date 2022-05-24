package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	service "github.com/my-epoch/api-gateway/gen/go/api/proto/v1"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"google.golang.org/grpc"
)

func registerServices(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	registerObjectService(mux, ctx, opts)
}

func registerObjectService(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	if err := service.RegisterObjectServiceHandlerFromEndpoint(ctx, mux, "object.service.consul", opts); err != nil {
		logger.Warnf("cannot register object service: %e", err)
		return
	}
}
