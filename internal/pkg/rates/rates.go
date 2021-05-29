package rates

import (
	"time"

	"github.com/goakshit/gandalf/internal/types"
)

const (
	TwoWheelerRatePerHour            float64 = 10
	FourWheelerRatePerHour           float64 = 20
	HeavyRatePerHour                 float64 = 40
	ChargesIfDurationLessThanOneHour float64 = 10
)

type RateService interface {
	GetRatesByType(duration time.Time, vehicleType string) (float64, error)
}

type service struct{}

func NewRatesService() *service {
	return &service{}
}

func (s *service) GetRatesByType(vehicleType string, arrivalTime, departureTime time.Time) (float64, error) {

	if vehicleType == "" {
		return 0, types.ErrRatesMissingVehicleType
	}

	if arrivalTime.IsZero() || arrivalTime.After(departureTime) {
		return 0, types.ErrRatesInvalidDate
	}

	var duration float64
	if departureTime.IsZero() {
		// Consider the usecase when vehicle has not left parking lot
		// Set departure time as current time
		departureTime = time.Now().UTC()
	}
	duration = departureTime.Sub(arrivalTime.UTC()).Hours()

	// If car was parked for less than 1hr, minimum charges will incur.
	if duration == 0 {
		return ChargesIfDurationLessThanOneHour, nil
	}

	// vehicle types include three: two, four & heavy
	switch vehicleType {
	case "two":
		return TwoWheelerRatePerHour * duration, nil
	case "heavy":
		return HeavyRatePerHour * duration, nil
	default:
		return FourWheelerRatePerHour * duration, nil // Consider default as "four"
	}
}
