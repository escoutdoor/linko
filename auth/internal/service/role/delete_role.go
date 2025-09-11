package role

import (
	"context"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) DeleteRole(ctx context.Context, roleID string) error {
	_, err := s.roleRepository.GetRole(ctx, roleID)
	if err != nil {
		return errwrap.Wrap("get role from repository", err)
	}

	if err := s.roleRepository.DeleteRole(ctx, roleID); err != nil {
		return errwrap.Wrap("delete role", err)
	}

	return nil
}
