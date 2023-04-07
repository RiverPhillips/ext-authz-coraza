// main is the entry point of the application.
package main

import (
	"github.com/RiverPhillips/ext-authz-coraza/grpcserver"
	"github.com/RiverPhillips/ext-authz-coraza/healthcheck"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	fx.New(
		fx.Provide(
			grpcserver.NewGrpcServer,
			grpcserver.AsService(healthcheck.NewHealthcheckService),
			zap.NewProduction,
		),
		fx.Invoke(func(srv *grpc.Server) {}),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
