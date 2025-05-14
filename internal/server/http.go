package server

import (
	"log"
	"net/http"

	handler "github.com/inidaname/mosque/auth_service/internal/handler/auth"
	"github.com/inidaname/mosque/auth_service/internal/service"
	"github.com/inidaname/mosque/auth_service/internal/types"
)

type httpServer struct {
	app types.Application
}

func NewHttpServer(app *types.Application) *httpServer {
	return &httpServer{app: *app}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	authService := service.NewAuthService(&s.app)
	authHandler := handler.NewHttpAuthHandler(*authService)
	authHandler.RegisterRouter(router)

	log.Println("Starting server on", ":"+s.app.Config.Server.HTTPPort)

	return http.ListenAndServe(":"+s.app.Config.Server.HTTPPort, router)
}
