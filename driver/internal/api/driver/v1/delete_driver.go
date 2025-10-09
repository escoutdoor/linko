package v1

import (
	"context"

	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
)

func (a *api) DeleteDriver(ctx context.Context, req *driverv1.DeleteDriverRequest) (*driverv1.DeleteDriverResponse, error) {
	if err := a.driverService.DeleteDriver(ctx, req.GetDriverId()); err != nil {
		return nil, err
	}

	return &driverv1.DeleteDriverResponse{}, nil
}
