package types

type CreateVehicleDetailsRecordResponse struct {
	StatusCode int `json:"status_code"`
}

type GetVehicleDetailsRequest struct {
	ID string `json:"id"`
}

type GetVehicleParkingDurationResponse struct {
	DurationInHours float64 `json:"duration_in_hours,omitempty"`
}

type GetVehicleParkingCostResponse struct {
	Cost float64 `json:"cost,omitempty"`
}
