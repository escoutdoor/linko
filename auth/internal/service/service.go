package service

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
)

type UserService interface {
	GetUser(ctx context.Context, userID string) (entity.User, error)
	UpdateUser(ctx context.Context, in dto.UpdateUserParams) (entity.User, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (entity.TokenPair, error)
	Register(ctx context.Context, in dto.CreateUserParams) (entity.TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (entity.TokenPair, error)
}

type RoleService interface {
	CreateRole(ctx context.Context, name string) (entity.Role, error)
	GetRole(ctx context.Context, roleID string) (entity.Role, error)
	ListRoles(ctx context.Context, in dto.ListRolesParams) ([]entity.Role, error)
	DeleteRole(ctx context.Context, roleID string) error
}
