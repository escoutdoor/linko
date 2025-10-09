package v1

import (
	"context"

	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/api/converter"
)

func (a *api) ListDrivers(ctx context.Context, req *driverv1.ListDriversRequest) (*driverv1.ListDriversResponse, error) {
	drivers, err := a.driverService.ListDrivers(ctx, converter.ProtoListDriversRequestToListDriversParams(req))
	if err != nil {
		return nil, err
	}

	protoDrivers := converter.DriversToProtoDrivers(drivers)
	return &driverv1.ListDriversResponse{Drivers: protoDrivers}, nil
}
