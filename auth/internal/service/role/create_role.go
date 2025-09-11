package role

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/entity"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) CreateRole(ctx context.Context, name string) (entity.Role, error) {
	role, err := s.roleRepository.CreateRole(ctx, name)
	if err != nil {
		return entity.Role{}, errwrap.Wrap("create role", err)
	}

	return role, nil
}
