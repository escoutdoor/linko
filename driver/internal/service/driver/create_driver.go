package driver

import (
	"context"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
)

func (s *service) CreateDriver(ctx context.Context, in dto.CreateDriverParams) (entity.Driver, error) {
	driver, err := s.driverRepository.CreateDriver(ctx, in)
	if err != nil {
		return entity.Driver{}, errwrap.Wrap("create driver", err)
	}

	return driver, nil
}
