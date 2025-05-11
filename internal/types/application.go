package types

import (
	"log/slog"

	// cache "github.com/inidaname/mosque/auth_service/cache"
	"github.com/inidaname/mosque/auth_service/internal/cache"
	db "github.com/inidaname/mosque/auth_service/internal/db/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Config        Config
	Logger        *slog.Logger
	Store         *db.Queries
	Db            *pgxpool.Pool
	Authenticator Authenticator
	// Mailer        mailer.Mailer
	Cache               cache.CacheService
	HealthAuthenticator HealthCheckableAuthenticator
}
