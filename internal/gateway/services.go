package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/my-epoch/api-gateway/pkg/logger"
	objectService "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	storageService "github.com/my-epoch/storage_service/gen/go/api/proto/v1"
	"google.golang.org/grpc"
	"net"
)

func resolveServicePort(name string) int {
	_, srvs, err := net.LookupSRV(name, "tcp", "service.consul")
	if err != nil {
		logger.Warnf("cannot resolve %s service port, continuing with default 50050", name)
		return 50050
	}

	if len(srvs) < 1 {
		logger.Warnf("cannot resolve %s service port, continuing with default 50050", name)
		return 50050
	}

	return int(srvs[0].Port)
}

func registerServices(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	registerObjectService(mux, ctx, opts)
	registerStorageService(mux, ctx, opts)
}

func registerObjectService(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	port := resolveServicePort("object")
	addr := fmt.Sprintf("%s:%d", "object.service.consul", port)
	if err := objectService.RegisterObjectServiceAPIHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		logger.Warnf("cannot register object service: %e", err)
		return
	}
}

func registerStorageService(mux *runtime.ServeMux, ctx context.Context, opts []grpc.DialOption) {
	port := resolveServicePort("storage")
	addr := fmt.Sprintf("%s:%d", "storage.service.consul", port)
	if err := storageService.RegisterStorageServiceAPIHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		logger.Warnf("cannot register storage service: %e", err)
		return
	}
}
