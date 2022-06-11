package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	objectService "github.com/my-epoch/object_service/gen/go/api/proto/v1"

	"github.com/my-epoch/api-gateway/pkg/logger"
	"google.golang.org/grpc"
)

func registerServices(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	registerObjectService(mux, ctx, opts)
}

func registerObjectService(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	if err := objectService.RegisterObjectServiceAPIHandlerFromEndpoint(ctx, mux, "object.service.consul:50050", opts); err != nil {
		logger.Warnf("cannot register object service: %e", err)
		return
	}
}
