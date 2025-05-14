package handler

import (
	"net/http"

	"github.com/inidaname/mosque/auth_service/internal/service"
	"github.com/inidaname/mosque/auth_service/internal/util"
	"github.com/inidaname/mosque_location/protos"
)

type AuthHttpHandler struct {
	authService service.AuthService
}

func NewHttpAuthHandler(authService service.AuthService) *AuthHttpHandler {
	handler := &AuthHttpHandler{
		authService: authService,
	}

	return handler
}

func (h *AuthHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /register", h.RegisterUser)
	router.HandleFunc("POST /login", h.LoginUser)
	router.HandleFunc("GET /health", h.Health)
}

func (h *AuthHttpHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req protos.RegisterUserRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.authService.RegisterUser(r.Context(), &req)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, res)
}

func (h *AuthHttpHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req protos.LoginUserRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.authService.LoginUser(r.Context(), &req)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, res)
}

func (h *AuthHttpHandler) Health(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, http.StatusOK, "")
}
