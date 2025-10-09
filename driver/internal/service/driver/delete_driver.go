package driver

import (
	"context"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

func (s *service) DeleteDriver(ctx context.Context, driverID string) error {
	_, err := s.driverRepository.GetDriver(ctx, driverID)
	if err != nil {
		return errwrap.Wrap("get driver from repository", err)
	}

	if err := s.driverRepository.DeleteDriver(ctx, driverID); err != nil {
		return errwrap.Wrap("delete driver", err)
	}
	return nil
}
