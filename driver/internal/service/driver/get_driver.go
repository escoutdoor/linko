package driver

import (
	"context"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/escoutdoor/linko/driver/internal/entity"
)

func (s *service) GetDriver(ctx context.Context, driverID string) (entity.Driver, error) {
	driver, err := s.driverRepository.GetDriver(ctx, driverID)
	if err != nil {
		return entity.Driver{}, errwrap.Wrap("get driver from repository", err)
	}

	return driver, nil
}
