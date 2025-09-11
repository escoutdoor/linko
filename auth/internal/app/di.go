package app

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/config"
	"github.com/escoutdoor/linko/auth/internal/repository"
	role_repository "github.com/escoutdoor/linko/auth/internal/repository/role"
	user_repository "github.com/escoutdoor/linko/auth/internal/repository/user"
	user_role_repository "github.com/escoutdoor/linko/auth/internal/repository/user_role"
	"github.com/escoutdoor/linko/auth/internal/service"
	auth_service "github.com/escoutdoor/linko/auth/internal/service/auth"
	role_service "github.com/escoutdoor/linko/auth/internal/service/role"
	user_service "github.com/escoutdoor/linko/auth/internal/service/user"
	"github.com/escoutdoor/linko/auth/internal/utils/token"
	"github.com/escoutdoor/linko/common/pkg/closer"
	"github.com/escoutdoor/linko/common/pkg/database"
	"github.com/escoutdoor/linko/common/pkg/database/pg"
	"github.com/escoutdoor/linko/common/pkg/database/txmanager"
	"github.com/escoutdoor/linko/common/pkg/logger"
)

type di struct {
	dbClient      database.Client
	tokenProvider token.Provider
	txManager     database.TxManager

	authService service.AuthService
	userService service.UserService
	roleService service.RoleService

	userRepository     repository.UserRepository
	roleRepository     repository.RoleRepository
	userRoleRepository repository.UserRoleRepository
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

func (d *di) RoleRepository(ctx context.Context) repository.RoleRepository {
	if d.roleRepository == nil {
		d.roleRepository = role_repository.NewRoleRepository(d.DBClient(ctx))
	}

	return d.roleRepository
}

func (d *di) UserRoleRepository(ctx context.Context) repository.UserRoleRepository {
	if d.userRoleRepository == nil {
		d.userRoleRepository = user_role_repository.NewUserRoleRepository(d.DBClient(ctx))
	}

	return d.userRoleRepository
}

func (d *di) RoleService(ctx context.Context) service.RoleService {
	if d.roleService == nil {
		d.roleService = role_service.NewService(d.RoleRepository(ctx))
	}

	return d.roleService
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

func (d *di) TxManager(ctx context.Context) database.TxManager {
	if d.txManager == nil {
		d.txManager = txmanager.NewTransactionManager(d.DBClient(ctx).DB())
	}

	return d.txManager
}

func (d *di) UserService(ctx context.Context) service.UserService {
	if d.userService == nil {
		d.userService = user_service.NewService(
			d.UserRepository(ctx),
			d.RoleRepository(ctx),
			d.UserRoleRepository(ctx),
			d.TxManager(ctx),
		)
	}

	return d.userService
}

func (d *di) AuthService(ctx context.Context) service.AuthService {
	if d.authService == nil {
		d.authService = auth_service.NewService(
			d.UserRepository(ctx),
			d.RoleRepository(ctx),
			d.UserRoleRepository(ctx),
			d.TokenProvider(ctx),
			d.TxManager(ctx),
		)
	}

	return d.authService
}
