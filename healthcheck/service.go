// Package healthcheck implements the gRPC health check service.
package healthcheck

import (
	"github.com/RiverPhillips/ext-authz-coraza/grpcserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// This ensures that healthService implements grpcserver.Service
var _ grpcserver.Service = (*healthService)(nil)

type healthService struct {
	server      *health.Server
	description *grpc.ServiceDesc
}

// Register registers the health check service with the gRPC server.
func (h *healthService) Register(server *grpc.Server) {
	server.RegisterService(h.description, h.server)
}

// NewHealthcheckService returns a new healthcheck service.
func NewHealthcheckService() grpcserver.Service {
	srv := health.NewServer()
	srv.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	return &healthService{
		description: &grpc_health_v1.Health_ServiceDesc,
		server:      srv,
	}
}
