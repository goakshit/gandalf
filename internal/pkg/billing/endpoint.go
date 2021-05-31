package billing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/goakshit/gandalf/internal/types"
)

// //the endpoint will receive a request, convert to the desired
// //format, invoke the service and return the response structure
// func CreateVehicleDetailsRecordsEndpoint(svc billing.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(types.VehicleDetails)
// 		err := svc.CreateVehicleParkingRecord(ctx, req)
// 		if err != nil {
// 			return types.CreateVehicleDetailsRecordResponse{StatusCode: http.StatusInternalServerError}, err
// 		}
// 		return types.CreateVehicleDetailsRecordResponse{StatusCode: http.StatusOK}, err
// 	}
// }

//the endpoint will receive a request, convert to the desired
//format, invoke the service and return the response structure
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

//the endpoint will receive a request, convert to the desired
//format, invoke the service and return the response structure
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
