package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/rashad-j/device-management-api/internal/core/domain"
	"github.com/rashad-j/device-management-api/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func TestDeviceService_AddDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)

	service := usecases.NewDeviceService(mockRepo)

	device := domain.CreateDeviceInput{
		Name:  "iPhone 12",
		Brand: "Apple",
	}

	mockRepo.EXPECT().Create(gomock.Any()).Times(1)

	result := service.CreateDevice(device)

	assert.Equal(t, device.Name, result.Name)
	assert.Equal(t, device.Brand, result.Brand)
}

func TestDeviceService_GetDeviceByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)

	service := usecases.NewDeviceService(mockRepo)

	id := uuid.New()
	expectedDevice := domain.Device{
		ID:    id,
		Name:  "iPhone 12",
		Brand: "Apple",
	}

	mockRepo.EXPECT().GetById(id).Return(expectedDevice, nil).Times(1)

	result, err := service.GetDeviceById(id)

	assert.Nil(t, err)
	assert.Equal(t, expectedDevice, result)
}

func TestDeviceService_DeleteDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)
	service := usecases.NewDeviceService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil).Times(1)

	err := service.DeleteDevice(id)

	assert.Nil(t, err)
}

func TestDeviceService_UpdateDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)
	service := usecases.NewDeviceService(mockRepo)

	id := uuid.New()
	device := domain.CreateDeviceInput{
		Name:  "iPhone 13",
		Brand: "Apple",
	}

	mockRepo.EXPECT().Update(id, device).Return(nil).Times(1)

	result, err := service.UpdateDevice(id, device)
	assert.Equal(t, device.Name, result.Name)
	assert.Nil(t, err)
}

func TestDeviceService_ListDevices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)
	service := usecases.NewDeviceService(mockRepo)

	devices := []domain.Device{
		{ID: uuid.New(), Name: "iPhone 12", Brand: "Apple"},
		{ID: uuid.New(), Name: "Galaxy S21", Brand: "Samsung"},
	}

	mockRepo.EXPECT().List().Return(devices).Times(1)

	result := service.ListDevices()
	assert.Equal(t, len(devices), len(result))
	assert.Equal(t, devices, result)
}

func TestDeviceService_SearchByBrand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)
	service := usecases.NewDeviceService(mockRepo)

	brand := "Apple"
	devices := []domain.Device{
		{ID: uuid.New(), Name: "iPhone 12", Brand: "Apple"},
	}

	mockRepo.EXPECT().SearchByBrand(brand).Return(devices).Times(1)

	result := service.SearchDevicesByBrand(brand)
	assert.Equal(t, len(devices), len(result))
}
