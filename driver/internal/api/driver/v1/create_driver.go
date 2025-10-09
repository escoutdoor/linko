package v1

import (
	"context"

	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/api/converter"
)

func (a *api) CreateDriver(ctx context.Context, req *driverv1.CreateDriverRequest) (*driverv1.CreateDriverResponse, error) {
	driver, err := a.driverService.CreateDriver(ctx, converter.ProtoCreateDriverRequestToCreateDriverParams(req))
	if err != nil {
		return nil, err
	}

	protoDriver := converter.DriverToProtoDriver(driver)
	return &driverv1.CreateDriverResponse{Driver: protoDriver}, nil
}
