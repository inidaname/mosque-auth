package handler

import (
	"context"

	"github.com/inidaname/mosque/auth_service/internal/service"
	"github.com/inidaname/mosque_location/protos"
	"google.golang.org/grpc"
)

type AuthGrpcHandler struct {
	authService service.AuthService
	protos.UnimplementedAuthServiceServer
}

func NewGrpcAuthService(grpc *grpc.Server, authService service.AuthService) {
	grpcHandler := &AuthGrpcHandler{
		authService: authService,
	}

	protos.RegisterAuthServiceServer(grpc, grpcHandler)

}

func (h *AuthGrpcHandler) RegisterUser(ctx context.Context, req *protos.RegisterUserRequest) (*protos.RegisterUserResponse, error) {
	resp, err := h.authService.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *AuthGrpcHandler) LoginUser(ctx context.Context, req *protos.LoginUserRequest) (*protos.LoginUserResponse, error) {
	resp, err := h.authService.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
