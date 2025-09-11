package role

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/entity"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) GetRole(ctx context.Context, roleID string) (entity.Role, error) {
	role, err := s.roleRepository.GetRole(ctx, roleID)
	if err != nil {
		return entity.Role{}, errwrap.Wrap("get role from repository", err)
	}

	return role, nil
}
