package driver

import (
	"context"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
)

func (s *service) UpdateDriver(ctx context.Context, in dto.UpdateDriverParams) (entity.Driver, error) {
	_, err := s.driverRepository.GetDriver(ctx, in.DriverID)
	if err != nil {
		return entity.Driver{}, errwrap.Wrap("get driver from repository before update operation", err)
	}

	if err := s.driverRepository.UpdateDriver(ctx, in); err != nil {
		return entity.Driver{}, errwrap.Wrap("update driver", err)
	}

	driver, err := s.driverRepository.GetDriver(ctx, in.DriverID)
	if err != nil {
		return entity.Driver{}, errwrap.Wrap("get updated driver from repository", err)
	}

	return driver, nil
}
