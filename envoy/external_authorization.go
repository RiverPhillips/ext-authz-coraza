// Package envoy implements envoy's gRPC API for external authorization.
package envoy

import (
	context "context"

	"github.com/RiverPhillips/ext-authz-coraza/grpcserver"
	ext_authz_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
)

var _ grpcserver.Service = (*extAuthzServer)(nil)
var _ ext_authz_v3.AuthorizationServer = (*extAuthzServer)(nil)

type extAuthzServer struct{}

// Registers the extAuthzServer with the gRPC server.
func (e *extAuthzServer) Register(s *grpc.Server) {
	ext_authz_v3.RegisterAuthorizationServer(s, e)
}

// Check implements authv3.AuthorizationServer
func (*extAuthzServer) Check(context.Context, *ext_authz_v3.CheckRequest) (*ext_authz_v3.CheckResponse, error) {
	panic("unimplemented")
}

// NewExtAuthzService returns a new Ext AuthZ Service.
func NewExtAuthzService() grpcserver.Service {
	return &extAuthzServer{}
}
