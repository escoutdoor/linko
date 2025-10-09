package v1

import (
	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/service"
)

type api struct {
	driverService service.DriverService
	driverv1.UnimplementedDriverServiceServer
}

func New(driverService service.DriverService) *api {
	return &api{
		driverService: driverService,
	}
}
