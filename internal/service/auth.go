package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	db "github.com/inidaname/mosque/auth_service/internal/db/models"
	"github.com/inidaname/mosque/auth_service/internal/types"
	"github.com/inidaname/mosque/auth_service/internal/util"
	"github.com/inidaname/mosque/protos"
)

type AuthService struct {
	*types.Application
}

func NewAuthService(cfg *types.Application) *AuthService {
	return &AuthService{cfg}
}

func (s *AuthService) RegisterUser(ctx context.Context, req *protos.RegisterUserRequest) (*protos.RegisterUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.Store.CreateUser(ctx, db.CreateUserParams{
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    *req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &protos.RegisterUserResponse{
		User: &protos.User{
			Id:       user.ID.String(),
			Email:    user.Email,
			Phone:    user.Phone,
			FullName: user.FullName,
		},
	}, nil
}

func (s *AuthService) LoginUser(ctx context.Context, req *protos.LoginUserRequest) (*protos.LoginUserResponse, error) {
	user, err := s.Store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !util.VerifyPassword(user.Password, req.Password) {
		return nil, errors.New("Invalid Email or Password")
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(s.Config.Auth.Token.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": s.Config.Auth.Token.Iss,
		"aud": s.Config.Auth.Token.Iss,
	}

	token, err := s.Authenticator.GenerateToken(claims)

	if err != nil {
		s.Logger.Error("Unable to generate jwt token")
		return nil, err
	}

	return &protos.LoginUserResponse{
		Message: token,
		Success: true,
	}, nil
}
