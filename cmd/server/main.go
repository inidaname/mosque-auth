package main

import (
	"github.com/inidaname/mosque/auth_service/internal/config"
	"github.com/inidaname/mosque/auth_service/internal/server"
)

func main() {
	cfg := config.CreateApplication()

	httpServer := server.NewHttpServer(cfg)
	go httpServer.Run()

	grpcServer := server.NewGRPCServer(cfg)
	grpcServer.ListenAndServe()
}
