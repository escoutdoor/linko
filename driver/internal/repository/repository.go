package repository

import (
	"context"

	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
)

type DriverRepository interface {
	CreateDriver(ctx context.Context, in dto.CreateDriverParams) (entity.Driver, error)
	UpdateDriver(ctx context.Context, in dto.UpdateDriverParams) error
	GetDriver(ctx context.Context, driverID string) (entity.Driver, error)
	ListDrivers(ctx context.Context, in dto.ListDriversParams) ([]entity.Driver, error)
	DeleteDriver(ctx context.Context, driverID string) error
}
