package grpcserver

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// This ensures that service implements Service
var _ Service = (*service)(nil)

// Service is a gRPC service that can be registered with a gRPC server.
//
//go:generate mockery --name Service --quiet --output .
type Service interface {
	Register(*grpc.Server)
}

type service struct {
	desc   *grpc.ServiceDesc
	server interface{}
}

// Register implements Service
func (s *service) Register(srv *grpc.Server) {
	srv.RegisterService(s.desc, s.server)
}

// Desc implements Service
func (s *service) Desc() *grpc.ServiceDesc {
	return s.desc
}

// Server implements Service
func (s *service) Server() interface{} {
	return s.server
}

// NewService returns a new Service.
func NewService(desc *grpc.ServiceDesc, server interface{}) Service {
	return &service{
		desc:   desc,
		server: server,
	}
}

// AsService is a helper function to annotate a function as a gRPC service.
func AsService(f any) any {
	return fx.Annotate(f, fx.As(new(Service)), fx.ResultTags(`group:"grpc_service"`))
}
