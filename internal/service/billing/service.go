package billing

import (
	"context"
	"time"

	"github.com/goakshit/gandalf/internal/pkg/rates"
	"github.com/goakshit/gandalf/internal/types"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewBillingService(db *gorm.DB) *service {
	return &service{
		db: db,
	}
}

func (s *service) CreateVehicleParkingRecord(ctx context.Context, data types.VehicleDetails) error {
	return s.db.Table("vehicle_details").Create(data).Error
}

func (s *service) GetVehicleParkingDuration(ctx context.Context, ID string) (time.Duration, error) {

	var (
		duration       time.Duration
		vehicleDetails types.VehicleDetails
	)

	if ID == "" {
		return duration, types.ErrServiceBillingInvalidOrMissingID
	}
	err := s.db.Table("vehicle_details").Where("id = ?", ID).First(&vehicleDetails).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return duration, types.ErrServiceBillingRecordNotFound
		}
		return duration, types.ErrServiceBillingInternalServerError
	}
	return time.Now().UTC().Sub(vehicleDetails.ArrivedAt), nil
}

func (s *service) GetVehicleParkingCost(ctx context.Context, ID string) (float64, error) {

	var (
		cost           float64
		vehicleDetails types.VehicleDetails
	)
	if ID == "" {
		return cost, types.ErrServiceBillingInvalidOrMissingID
	}
	err := s.db.Table("vehicle_details").Where("id = ?", ID).First(&vehicleDetails).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return cost, types.ErrServiceBillingRecordNotFound
		}
		return cost, types.ErrServiceBillingInternalServerError
	}
	ratesSvc := rates.NewRatesService()
	return ratesSvc.GetRatesByType(vehicleDetails.RegNo, vehicleDetails.ArrivedAt, vehicleDetails.DepartedAt)
}
