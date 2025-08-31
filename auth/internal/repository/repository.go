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
