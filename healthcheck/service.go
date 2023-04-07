// Package healthcheck implements the gRPC health check service.
package healthcheck

import (
	"github.com/RiverPhillips/ext-authz-coraza/grpcserver"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// NewHealthcheckService returns a new healthcheck service.
func NewHealthcheckService() grpcserver.Service {
	srv := health.NewServer()
	srv.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	return grpcserver.NewService(
		&grpc_health_v1.Health_ServiceDesc,
		srv,
	)
}
