package server

import (
	"log"
	"net"

	handler "github.com/inidaname/mosque/auth_service/internal/handler/auth"
	"github.com/inidaname/mosque/auth_service/internal/service"
	"github.com/inidaname/mosque/auth_service/internal/types"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	app types.Application
}

func NewGRPCServer(app *types.Application) *GRPCServer {
	return &GRPCServer{app: *app}
}

func (s *GRPCServer) ListenAndServe() error {
	lis, err := net.Listen("tcp", ":"+s.app.Config.Server.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authService := service.NewAuthService(&s.app)
	handler.NewGrpcAuthService(grpcServer, *authService)
	// repo := &repository.PostgresUserRepo{Db: s.pool}
	// svc := &service.UserService{Repo: repo}
	// h := &handler.UserHandler{Service: svc}

	// pb.RegisterAuthServiceServer(grpcServer, h)
	log.Println("GRPC Auth Service running on :", s.app.Config.Server.GRPCPort)
	return grpcServer.Serve(lis)
}
