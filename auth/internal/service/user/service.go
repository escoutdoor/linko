package user

import (
	"github.com/escoutdoor/linko/auth/internal/repository"
	"github.com/escoutdoor/linko/common/pkg/database"
)

type service struct {
	userRepository     repository.UserRepository
	roleRepository     repository.RoleRepository
	userRoleRepository repository.UserRoleRepository

	txManager database.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	userRoleRepository repository.UserRoleRepository,
	txManager database.TxManager,
) *service {
	return &service{
		userRepository:     userRepository,
		roleRepository:     roleRepository,
		userRoleRepository: userRoleRepository,
		txManager:          txManager,
	}
}
