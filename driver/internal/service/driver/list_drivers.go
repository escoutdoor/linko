package driver

import (
	"context"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
)

func (s *service) ListDrivers(ctx context.Context, in dto.ListDriversParams) ([]entity.Driver, error) {
	limit := in.PageSize + 1
	drivers, err := s.driverRepository.ListDrivers(ctx, limit, in.Cursor)
	if err != nil {
		return nil, errwrap.Wrap("list drivers from repository", err)
	}

	return drivers, nil
}
