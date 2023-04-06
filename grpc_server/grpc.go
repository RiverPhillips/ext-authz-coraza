// grpc_server contains the gRPC server and its dependencies.
package grpc_server

import (
	"context"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServerParams struct {
	fx.In

	LC       fx.Lifecycle
	Services []Service `group:"grpc_service"`
	Log      *zap.Logger
}

func NewGrpcServer(p GrpcServerParams) *grpc.Server {
	srv := grpc.NewServer()

	for _, service := range p.Services {
		srv.RegisterService(service.Desc(), service.Server())
	}

	reflection.Register(srv)

	p.LC.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				port := ":8080"
				ln, err := net.Listen("tcp", port)
				if err != nil {
					return err
				}
				p.Log.Info("Starting gRPC server", zap.String("port", port))
				go srv.Serve(ln)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				srv.GracefulStop()
				p.Log.Info("Stopped gRPC server")
				return nil
			},
		},
	)

	return srv
}
