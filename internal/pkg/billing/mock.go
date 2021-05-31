package billing

import (
	"context"
	"time"

	"github.com/goakshit/gandalf/internal/types"
	"github.com/stretchr/testify/mock"
)

type billingMock struct {
	mock.Mock
}

// GetbillingMock - Returns service mock
// Mocks service interface
func GetbillingMock() *billingMock {
	return &billingMock{}
}

func (m *billingMock) CreateVehicleParkingRecord(ctx context.Context, data types.VehicleDetails) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *billingMock) GetVehicleParkingDuration(ctx context.Context, ID string) (time.Duration, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(time.Duration), args.Error(1)
}

func (m *billingMock) GetVehicleParkingCost(ctx context.Context, ID string) (float64, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(float64), args.Error(1)
}
