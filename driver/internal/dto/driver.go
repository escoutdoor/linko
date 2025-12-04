package dto

import "github.com/escoutdoor/linko/driver/internal/pagination"

type CreateDriverParams struct {
	UserID  string
	Vehicle *CreateVehicleParams
}

type UpdateDriverParams struct {
	DriverID string
	Status   *int32
	Vehicle  *UpdateVehicleParams
}

type ListDriversParams struct {
	PageSize int32
	Cursor   *pagination.Cursor
}

type CreateVehicleParams struct {
	Type        *int32
	Model       *string
	PlateNumber *string
	Color       *string
}

type UpdateVehicleParams struct {
	Type        int32
	Model       string
	PlateNumber string
	Color       string
}
