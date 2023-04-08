// Package grpcserver contains the gRPC server and its dependencies.
package grpcserver

import (
	"context"
	"net"

	ext_authz_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServerParams struct {
	fx.In

	LC       fx.Lifecycle
	Services []Service `group:"grpc_service"`
	Log      *zap.Logger
	ExtAuthz ext_authz_v3.AuthorizationServer
}

// NewGrpcServer returns a new gRPC server.
func NewGrpcServer(p grpcServerParams) *grpc.Server {
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
		srv.RegisterService(service.Desc(), service.Server())
	}

	reflection.Register(srv)

	ext_authz_v3.RegisterAuthorizationServer(srv, p.ExtAuthz)

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
