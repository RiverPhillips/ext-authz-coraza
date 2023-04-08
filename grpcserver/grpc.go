// Package grpcserver contains the gRPC server and its dependencies.
package grpcserver

import (
	"context"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGrpcServerParams is injected into NewGrpcServer by fx.
type NewGrpcServerParams struct {
	fx.In

	LC       fx.Lifecycle
	Services []Service `group:"grpc_service"`
	Log      *zap.Logger
}

// NewGrpcServer returns a new gRPC server.
func NewGrpcServer(p NewGrpcServerParams) *grpc.Server {
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(p.Log),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(p.Log),
		)),
	)

	for _, service := range p.Services {
		service.Register(srv)
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
				go func() {
					if err := srv.Serve(ln); err != nil {
						p.Log.Error("Failed to serve gRPC server", zap.Error(err))
					}
				}()
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
