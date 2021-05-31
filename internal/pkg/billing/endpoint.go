package billing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/goakshit/gandalf/internal/types"
)

// swagger:route GET /duration/{id} billing getVehicleParkingDuration
// Returns the duration of parked vehicle from id
// responses:
//	200: GetVehicleParkingDurationResponse
func getVehicleParkingDurationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.GetVehicleDetailsRequest)
		duration, err := svc.GetVehicleParkingDuration(ctx, req.ID)
		if err != nil {
			return types.GetVehicleParkingDurationResponse{}, err
		}
		return types.GetVehicleParkingDurationResponse{
			DurationInHours: duration.Hours(),
		}, err
	}
}

// swagger:route GET /cost/{id} billing getVehicleParkingCost
// Returns the cost of parking vehicle for duration by id
// responses:
//	200: GetVehicleParkingCostResponse
func getVehicleParkingCostEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.GetVehicleDetailsRequest)
		cost, err := svc.GetVehicleParkingCost(ctx, req.ID)
		if err != nil {
			return types.GetVehicleParkingCostResponse{}, err
		}
		return types.GetVehicleParkingCostResponse{
			Cost: cost,
		}, err
	}
}
