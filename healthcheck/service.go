// Healthcheck implements the gRPC health check service.
package healthcheck

import (
	"github.com/RiverPhillips/ext-authz-coraza/grpc_server"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewHealthcheckService() grpc_server.Service {
	srv := health.NewServer()
	srv.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	return grpc_server.NewService(
		&grpc_health_v1.Health_ServiceDesc,
		srv,
	)
}
