package dto

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
	PageSize  int32
	PageToken string
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
