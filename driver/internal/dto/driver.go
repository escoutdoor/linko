package dto

type CreateDriverParams struct {
	UserID  string
	Vehicle *VehicleParams
}

type UpdateDriverParams struct {
	DriverID string
	Status   *int32
	Vehicle  *VehicleParams
}

type ListDriversParams struct {
	PageSize  int32
	PageToken string
}

type VehicleParams struct {
	Type        int32
	Model       string
	PlateNumber string
	Color       string
}
