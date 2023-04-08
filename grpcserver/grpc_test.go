package grpcserver_test

import (
	"testing"

	"github.com/RiverPhillips/ext-authz-coraza/grpcserver"
	"github.com/stretchr/testify/mock"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type mockLifecycle struct{}

func (m mockLifecycle) Append(fx.Hook) {

}

func TestGrpcServerRegistersAllServices(t *testing.T) {
	mockService := grpcserver.NewMockService(t)
	mockService.Test(t)

	mockService.On("Register", mock.Anything).Return()

	params := grpcserver.NewGrpcServerParams{
		Log: zap.NewNop(),
		LC:  mockLifecycle{},
		Services: []grpcserver.Service{
			mockService,
		},
	}

	_ = grpcserver.NewGrpcServer(params)

	mockService.AssertNumberOfCalls(t, "Register", 1)
}
