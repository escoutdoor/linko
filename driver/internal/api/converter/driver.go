package converter

import (
	driverv1 "github.com/escoutdoor/linko/common/pkg/proto/driver/v1"
	"github.com/escoutdoor/linko/driver/internal/dto"
	"github.com/escoutdoor/linko/driver/internal/entity"
	apperrors "github.com/escoutdoor/linko/driver/internal/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	statusFieldName  = "status"
	vehicleFieldName = "vehicle"
)

func DriverToProtoDriver(driver entity.Driver) *driverv1.Driver {
	return &driverv1.Driver{
		Id:          driver.ID,
		UserId:      driver.UserID,
		Rating:      driver.Rating,
		ReviewCount: driver.ReviewCount,
		Vehicle: &driverv1.Vehicle{
			Type:        driverv1.VehicleType(driver.Vehicle.Type),
			Model:       driver.Vehicle.Model,
			PlateNumber: driver.Vehicle.PlateNumber,
			Color:       driver.Vehicle.Color,
		},
		CreatedAt: timestamppb.New(driver.CreatedAt),
		UpdatedAt: timestamppb.New(driver.UpdatedAt),
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
	params := dto.CreateDriverParams{UserID: req.GetUserId()}
	vehicle := req.GetVehicle()

	if req.GetVehicle() != nil {
		params.Vehicle = &dto.VehicleParams{
			Type:        int32(vehicle.GetType()),
			Model:       vehicle.GetModel(),
			PlateNumber: vehicle.GetPlateNumber(),
			Color:       vehicle.GetColor(),
		}
	}

	return params
}

func ProtoListDriversRequestToListDriversParams(req *driverv1.ListDriversRequest) dto.ListDriversParams {
	return dto.ListDriversParams{
		PageSize:  req.GetPageSize(),
		PageToken: req.GetPageToken(),
	}
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
		// TODO
		switch path {
		case statusFieldName:
			v := update.GetStatus()
			if v == 0 {
				return dto.UpdateDriverParams{}, apperrors.ValidationFailed("status is specified in update mask, but has no value")
			}
			status := int32(v)
			params.Status = &status
		case vehicleFieldName:
			v := update.GetVehicle()
			if v == nil {
				return dto.UpdateDriverParams{}, apperrors.ValidationFailed("vehicle is specified in update mask, but has no value")
			}

			vehicleType := v.GetType()
			if vehicleType == 0 {
				return dto.UpdateDriverParams{}, apperrors.ValidationFailed("vehicle type is specified in update mask, but has no value")
			}

			vehicle := &dto.VehicleParams{
				Type:        int32(vehicleType),
				Model:       v.GetModel(),
				PlateNumber: v.GetPlateNumber(),
				Color:       v.GetColor(),
			}
			params.Vehicle = vehicle
		}
	}

	return params, nil
}
