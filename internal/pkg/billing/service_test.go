package billing

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/goakshit/gandalf/internal/persistence"
	"github.com/goakshit/gandalf/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateVehicleParkingRecordSuccess(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewBillingService(mockRepo)
	err := svc.CreateVehicleParkingRecord(context.Background(), types.VehicleDetails{
		ID:        0,
		ArrivedAt: time.Now().UTC(),
		RegNo:     "JK02JY4439",
	})
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateVehicleParkingRecordFailure(t *testing.T) {

	e := errors.New("Internal Server Error")

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(e)

	svc := NewBillingService(mockRepo)
	err := svc.CreateVehicleParkingRecord(context.Background(), types.VehicleDetails{
		ID:        1,
		ArrivedAt: time.Now().UTC(),
		RegNo:     "JK02JY4439",
	})
	assert.EqualError(t, err, e.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleParkingDurationSuccess(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(nil)

	svc := NewBillingService(mockRepo)
	_, err := svc.GetVehicleParkingDuration(context.Background(), "0")
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleParkingDurationFailure(t *testing.T) {

	e := errors.New("Internal Server Error")

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(e)

	svc := NewBillingService(mockRepo)
	_, err := svc.GetVehicleParkingDuration(context.Background(), "0")
	assert.EqualError(t, err, types.ErrServiceBillingInternalServerError.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleParkingDurationRecordNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(gorm.ErrRecordNotFound)

	svc := NewBillingService(mockRepo)
	_, err := svc.GetVehicleParkingDuration(context.Background(), "0")
	assert.EqualError(t, err, types.ErrServiceBillingRecordNotFound.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleParkingDurationRecordBadRequest(t *testing.T) {

	mockRepo := persistence.GetGormMock()

	svc := NewBillingService(mockRepo)
	_, err := svc.GetVehicleParkingDuration(context.Background(), "")
	assert.EqualError(t, err, types.ErrServiceBillingInvalidOrMissingID.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleParkingCostRecordNotFound(t *testing.T) {

	mockRepo := persistence.GetGormMock()
	mockRepo.On("Table", mock.Anything).Return(mockRepo)
	mockRepo.On("Where", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("First", mock.Anything, mock.Anything).Return(mockRepo)
	mockRepo.On("Error").Return(gorm.ErrRecordNotFound)

	svc := NewBillingService(mockRepo)
	_, err := svc.GetVehicleParkingCost(context.Background(), "0")
	assert.EqualError(t, err, types.ErrServiceBillingRecordNotFound.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleParkingCostRecordBadRequest(t *testing.T) {

	mockRepo := persistence.GetGormMock()

	svc := NewBillingService(mockRepo)
	_, err := svc.GetVehicleParkingCost(context.Background(), "")
	assert.EqualError(t, err, types.ErrServiceBillingInvalidOrMissingID.Error())
	mockRepo.AssertExpectations(t)
}
