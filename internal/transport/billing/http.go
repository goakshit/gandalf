package billing

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	billingEpt "github.com/goakshit/gandalf/internal/endpoints/billing"
	billingSvc "github.com/goakshit/gandalf/internal/service/billing"
	"github.com/goakshit/gandalf/internal/types"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc billingSvc.Service, logger log.Logger) *mux.Router {
	//options provided by the Go kit to facilitate error control
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}
	//definition of a handler
	getParkingDurationHandler := httptransport.NewServer(
		billingEpt.GetVehicleParkingDurationEndpoint(svc),
		decodeBillingRequest, //converts the parameters received via the request body into the struct expected by the endpoint
		encodeResponse,       //converts the struct returned by the endpoint to a json response
		options...,
	)

	//definition of a handler
	getParkingCostHandler := httptransport.NewServer(
		billingEpt.GetVehicleParkingCostEndpoint(svc),
		decodeBillingRequest, //converts the parameters received via the request body into the struct expected by the endpoint
		encodeResponse,       //converts the struct returned by the endpoint to a json response
		options...,
	)
	r := mux.NewRouter() //I'm using Gorilla Mux, but it could be any other library, or even the stdlib
	r.Methods("GET").Path("/api/duration/{id}").Handler(getParkingDurationHandler)
	r.Methods("GET").Path("/api/cost/{id}").Handler(getParkingCostHandler)
	return r
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case types.ErrServiceBillingRecordNotFound:
		return http.StatusNotFound
	case types.ErrServiceBillingInvalidOrMissingID:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

//converts the parameters received via the request body into the struct expected by the endpoint
func decodeBillingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request types.GetVehicleDetailsRequest
	vars := mux.Vars(r)
	request.ID = vars["id"]
	return request, nil
}

//converts the struct returned by the endpoint to a json response
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
