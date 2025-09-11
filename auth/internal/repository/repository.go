package repository

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, in dto.CreateUserParams) (string, error)
	UpdateUser(ctx context.Context, in dto.UpdateUserParams) error
	GetUserByID(ctx context.Context, userID string) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type RoleRepository interface {
	CreateRole(ctx context.Context, name string) (entity.Role, error)
	UpdateRole(ctx context.Context, in dto.UpdateRoleParams) error
	DeleteRole(ctx context.Context, roleID string) error
	GetRole(ctx context.Context, roleID string) (entity.Role, error)
	ListRoles(ctx context.Context, in dto.ListRolesParams) ([]entity.Role, error)
}

type UserRoleRepository interface {
	AddRolesToUser(ctx context.Context, userID string, roles ...string) error
	ListUserRoles(ctx context.Context, userID string) ([]string, error)
}
