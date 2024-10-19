package usecases

import (
	"time"

	"github.com/rashad-j/device-management-api/internal/core/domain"
	"github.com/rashad-j/device-management-api/internal/core/ports"

	"github.com/google/uuid"
)

type DeviceService struct {
	repo ports.DeviceRepository
}

func NewDeviceService(repo ports.DeviceRepository) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) CreateDevice(input domain.CreateDeviceInput) domain.Device {
	device := domain.Device{
		ID:        uuid.New(),
		Name:      input.Name,
		Brand:     input.Brand,
		CreatedAt: time.Now(),
	}
	s.repo.Create(device)
	return device
}

func (s *DeviceService) GetDeviceById(id uuid.UUID) (domain.Device, error) {
	return s.repo.GetById(id)
}

func (s *DeviceService) ListDevices() []domain.Device {
	return s.repo.List()
}

func (s *DeviceService) UpdateDevice(id uuid.UUID, input domain.CreateDeviceInput) (domain.Device, error) {
	device, err := s.repo.GetById(id)
	if err != nil {
		return domain.Device{}, err
	}

	device.Name = input.Name
	device.Brand = input.Brand
	s.repo.Update(id, device)
	return device, nil
}

func (s *DeviceService) DeleteDevice(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *DeviceService) SearchDevicesByBrand(brand string) []domain.Device {
	return s.repo.SearchByBrand(brand)
}
