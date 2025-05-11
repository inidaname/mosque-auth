package config

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/inidaname/mosque/auth_service/internal/cache"
	db "github.com/inidaname/mosque/auth_service/internal/db"
	"github.com/inidaname/mosque/auth_service/internal/types"
	"github.com/inidaname/mosque/auth_service/internal/util"
)

var (
	instance *types.Application
	once     sync.Once
)

func CreateApplication() *types.Application {
	once.Do(func() {
		cfg, err := LoadConfig("internal/config/config.yaml")

		jwtAuthenticator := util.NewJWTAuthenticator(
			cfg.Auth.Token.Secret,
			cfg.Auth.Token.Iss,
			cfg.Auth.Token.Iss,
		)

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		store, dbPool, err := db.ConnectDB(ctx, logger, cfg)
		if err != nil {
			logger.Error("failed to connect to database", "error", err)
			os.Exit(1)
		}

		// Initialize thread-safe cache
		cacheService := cache.NewCacheService(5*time.Minute, 10*time.Minute)
		// Wrap the authenticator with health check capabilities
		healthAuth := util.NewHealthAuthenticator(
			jwtAuthenticator,
			logger.With("component", "auth_health"),
			15*time.Second, // Custom check interval
		)

		instance = &types.Application{
			Config:              *cfg,
			Logger:              logger,
			Store:               store,
			Db:                  dbPool,
			Cache:               *cacheService,
			Authenticator:       jwtAuthenticator,
			HealthAuthenticator: healthAuth,
		}
	})

	return instance
}
