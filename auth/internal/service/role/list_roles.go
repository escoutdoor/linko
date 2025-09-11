package role

import (
	"context"

	"github.com/escoutdoor/linko/auth/internal/dto"
	"github.com/escoutdoor/linko/auth/internal/entity"
	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) ListRoles(ctx context.Context, in dto.ListRolesParams) ([]entity.Role, error) {
	roles, err := s.roleRepository.ListRoles(ctx, in)
	if err != nil {
		return nil, errwrap.Wrap("get role from repository", err)
	}

	return roles, nil
}
