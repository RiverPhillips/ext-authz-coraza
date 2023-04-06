package grpc_server

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// This ensures that service implements Service
var _ Service = (*service)(nil)

type Service interface {
	Desc() *grpc.ServiceDesc
	Server() interface{}
}

type service struct {
	desc   *grpc.ServiceDesc
	server interface{}
}

// Desc implements Service
func (s *service) Desc() *grpc.ServiceDesc {
	return s.desc
}

// Server implements Service
func (s *service) Server() interface{} {
	return s.server
}

func NewService(desc *grpc.ServiceDesc, server interface{}) Service {
	return &service{
		desc:   desc,
		server: server,
	}
}

func AsService(f any) any {
	return fx.Annotate(f, fx.As(new(Service)), fx.ResultTags(`group:"grpc_service"`))
}
