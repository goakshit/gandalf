package billing

import (
	"context"
	"time"

	"github.com/goakshit/gandalf/internal/types"
)

type Service interface {
	CreateVehicleParkingRecord(ctx context.Context, data types.VehicleDetails) error
	GetVehicleParkingDuration(ctx context.Context, ID string) (time.Duration, error)
	GetVehicleParkingCost(ctx context.Context, ID string) (float64, error)
}
