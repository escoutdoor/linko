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
		return entity.Driver{}, errwrap.Wrap("get driver from repository", err)
	}

	driver, err := s.driverRepository.UpdateDriver(ctx, in)
	if err != nil {
		return entity.Driver{}, errwrap.Wrap("update driver", err)
	}

	return driver, nil
}
