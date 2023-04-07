// Package envoy implements envoy's gRPC API for external authorization.
package envoy

import (
	context "context"

	ext_authz_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
)

type extAuthzServer struct{}

// Check implements authv3.AuthorizationServer
func (*extAuthzServer) Check(context.Context, *ext_authz_v3.CheckRequest) (*ext_authz_v3.CheckResponse, error) {
	panic("unimplemented")
}

var _ ext_authz_v3.AuthorizationServer = (*extAuthzServer)(nil)
