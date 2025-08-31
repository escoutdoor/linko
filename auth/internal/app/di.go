package app

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/config"
	"github.com/escoutdoor/linko/auth/internal/repository"
	user_repository "github.com/escoutdoor/linko/auth/internal/repository/user"
	"github.com/escoutdoor/linko/auth/internal/service"
	auth_service "github.com/escoutdoor/linko/auth/internal/service/auth"
	user_service "github.com/escoutdoor/linko/auth/internal/service/user"
	"github.com/escoutdoor/linko/auth/internal/utils/token"
	"github.com/escoutdoor/linko/common/pkg/closer"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/escoutdoor/linko/common/pkg/database/pg"
	"github.com/escoutdoor/linko/common/pkg/logger"
)

type di struct {
	dbClient      database.Client
	tokenProvider token.Provider

	authService    service.AuthService
	userService    service.UserService
	userRepository repository.UserRepository
}

func newDiContainer() *di {
	return &di{}
}

func (d *di) DBClient(ctx context.Context) database.Client {
	if d.dbClient == nil {
		client, err := pg.NewClient(ctx, config.AppConfig().Postgres.Dsn())
		if err != nil {
			logger.Fatal(ctx, "new database client", err)
		}

		if err := client.DB().Ping(ctx); err != nil {
			logger.Fatal(ctx, "ping database: %s", err)
		}

		d.dbClient = client
		closer.Add(func(ctx context.Context) error {
			client.Close()
			return nil
		})
	}

	return d.dbClient
}

func (d *di) UserRepository(ctx context.Context) repository.UserRepository {
	if d.userRepository == nil {
		d.userRepository = user_repository.NewRepository(d.DBClient(ctx))
	}

	return d.userRepository
}

func (d *di) TokenProvider(ctx context.Context) token.Provider {
	if d.tokenProvider == nil {
		d.tokenProvider = token.NewTokenProvider(
			config.AppConfig().JwtToken.AccessTokenSecretKey(),
			config.AppConfig().JwtToken.RefreshTokenSecretKey(),
			config.AppConfig().JwtToken.AccessTokenTTL(),
			config.AppConfig().JwtToken.RefreshTokenTTL(),
		)
	}

	return d.tokenProvider
}

func (d *di) UserService(ctx context.Context) service.UserService {
	if d.userService == nil {
		d.userService = user_service.NewService(d.UserRepository(ctx))
	}

	return d.userService
}

func (d *di) AuthService(ctx context.Context) service.AuthService {
	if d.authService == nil {
		d.authService = auth_service.NewService(d.UserRepository(ctx), d.TokenProvider(ctx))
	}

	return d.authService
}
