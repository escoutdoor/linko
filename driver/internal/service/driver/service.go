package driver

import (
	"github.com/escoutdoor/linko/driver/internal/repository"
	def "github.com/escoutdoor/linko/driver/internal/service"
)

type service struct {
	driverRepository repository.DriverRepository
}

var _ def.DriverService = (*service)(nil)

func New(driverRepository repository.DriverRepository) *service {
	return &service{
		driverRepository: driverRepository,
	}
}
