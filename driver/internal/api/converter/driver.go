package converter

import (
	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
	apperrors "github.com/escoutdoor/linko/driver/internal/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DriverToProtoDriver(driver entity.Driver) *driverv1.Driver {
	rating := driver.TotalRatingSum / float64(driver.ReviewCount)

	return &driverv1.Driver{
		Id:          driver.ID,
		UserId:      driver.UserID,
		Rating:      float32(rating),
		ReviewCount: driver.ReviewCount,
		CreatedAt:   timestamppb.New(driver.CreatedAt),
		UpdatedAt:   timestamppb.New(driver.UpdatedAt),
	}
}

func DriversToProtoDrivers(drivers []entity.Driver) []*driverv1.Driver {
	list := make([]*driverv1.Driver, 0, len(drivers))
	for _, d := range drivers {
		list = append(list, DriverToProtoDriver(d))
	}

	return list
}

func ProtoCreateDriverRequestToCreateDriverParams(req *driverv1.CreateDriverRequest) dto.CreateDriverParams {
	return dto.CreateDriverParams{
		UserID: req.GetUserId(),
	}
}

func ProtoListDriversRequestToListDriversParams(req *driverv1.ListDriversRequest) dto.ListDriversParams {
	return dto.ListDriversParams{}
}

func ProtoUpdateDriverRequestToUpdateDriverParams(req *driverv1.UpdateDriverRequest) (dto.UpdateDriverParams, error) {
	params := dto.UpdateDriverParams{
		DriverID: req.GetDriverId(),
	}

	mask := req.GetUpdateMask()
	if mask == nil || len(mask.GetPaths()) == 0 {
		return dto.UpdateDriverParams{}, apperrors.ValidationFailed("update mask is not provided or empty")
	}

	update := req.GetUpdate()
	if update == nil {
		return dto.UpdateDriverParams{}, apperrors.ValidationFailed("update body is missing while update mask is provided")
	}

	for _, path := range mask.GetPaths() {
		switch path {
		}
	}

	return params, nil
}
