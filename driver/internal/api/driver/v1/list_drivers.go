package v1

import (
	"context"

	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/api/converter"
	"github.com/escoutdoor/linko/driver/internal/pagination"
)

func (a *api) ListDrivers(ctx context.Context, req *driverv1.ListDriversRequest) (*driverv1.ListDriversResponse, error) {
	in, err := converter.ProtoListDriversRequestToListDriversParams(req)
	if err != nil {
		return nil, err
	}

	drivers, err := a.driverService.ListDrivers(ctx, in)
	if err != nil {
		return nil, err
	}

	pageSize := int(in.PageSize)
	var nextPageToken string
	if len(drivers) > pageSize {
		lastItem := drivers[pageSize]
		nextPageToken = pagination.EncodeCursor(lastItem.ID, lastItem.CreatedAt)
		drivers = drivers[:pageSize]
	}

	protoDrivers := converter.DriversToProtoDrivers(drivers)
	return &driverv1.ListDriversResponse{
		Drivers:       protoDrivers,
		NextPageToken: nextPageToken,
	}, nil
}
