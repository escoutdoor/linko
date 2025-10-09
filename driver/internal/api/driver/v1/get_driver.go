package v1

import (
	"context"

	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/api/converter"
)

func (a *api) GetDriver(ctx context.Context, req *driverv1.GetDriverRequest) (*driverv1.GetDriverResponse, error) {
	driver, err := a.driverService.GetDriver(ctx, req.GetDriverId())
	if err != nil {
		return nil, err
	}

	protoDriver := converter.DriverToProtoDriver(driver)
	return &driverv1.GetDriverResponse{Driver: protoDriver}, nil
}
