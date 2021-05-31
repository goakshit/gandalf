package types

type CreateVehicleDetailsRecordResponse struct {
	StatusCode int `json:"status_code"`
}

type GetVehicleDetailsRequest struct {
	ID string `json:"id"`
}

// swagger:response
type GetVehicleParkingDurationResponse struct {
	// Duration(Hours) for which vehicle was parked
	DurationInHours float64 `json:"duration_in_hours,omitempty"`
}

// swagger:response
type GetVehicleParkingCostResponse struct {
	// Cost of parking vehicle for 'X' duration
	Cost float64 `json:"cost,omitempty"`
}
