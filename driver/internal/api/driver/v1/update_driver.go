package v1

import (
	"context"

	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
)

func (a *api) UpdateDriver(ctx context.Context, req *driverv1.UpdateDriverRequest) (*driverv1.UpdateDriverResponse, error) {
	return &driverv1.UpdateDriverResponse{}, nil
}
