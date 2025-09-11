package role

import (
	"github.com/escoutdoor/linko/auth/internal/repository"
)

type service struct {
	roleRepository repository.RoleRepository
}

func NewService(roleRepository repository.RoleRepository) *service {
	return &service{
		roleRepository: roleRepository,
	}
}
