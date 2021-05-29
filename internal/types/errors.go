package types

import "fmt"

var (
	ErrRatesMissingVehicleType           error = fmt.Errorf("Missing vehicle type")
	ErrRatesInvalidDate                  error = fmt.Errorf("Invalid/Empty date")
	ErrServiceBillingRecordNotFound      error = fmt.Errorf("Record with given id doesn't exist")
	ErrServiceBillingInternalServerError error = fmt.Errorf("Something went wrong fetching errors")
)
