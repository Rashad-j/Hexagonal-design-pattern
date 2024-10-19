package repository

import (
	"errors"
	"sync"

	"github.com/rashad-j/device-management-api/internal/core/domain"
	"github.com/rashad-j/device-management-api/internal/core/ports"

	"github.com/google/uuid"
)

type MemoryDeviceRepository struct {
	devices map[uuid.UUID]domain.Device
	mu      sync.Mutex
}

func NewMemoryDeviceRepository() ports.DeviceRepository {
	return &MemoryDeviceRepository{
		devices: make(map[uuid.UUID]domain.Device),
	}
}

func (r *MemoryDeviceRepository) Create(device domain.Device) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.devices[device.ID] = device
}

func (r *MemoryDeviceRepository) GetById(id uuid.UUID) (domain.Device, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	device, exists := r.devices[id]
	if !exists {
		return domain.Device{}, errors.New("device not found")
	}
	return device, nil
}

func (r *MemoryDeviceRepository) List() []domain.Device {
	r.mu.Lock()
	defer r.mu.Unlock()
	devices := make([]domain.Device, 0, len(r.devices))
	for _, device := range r.devices {
		devices = append(devices, device)
	}
	return devices
}

func (r *MemoryDeviceRepository) Update(id uuid.UUID, updatedDevice domain.Device) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.devices[id]; !exists {
		return errors.New("device not found")
	}
	r.devices[id] = updatedDevice
	return nil
}

func (r *MemoryDeviceRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.devices[id]; !exists {
		return errors.New("device not found")
	}
	delete(r.devices, id)
	return nil
}

func (r *MemoryDeviceRepository) SearchByBrand(brand string) []domain.Device {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []domain.Device
	for _, device := range r.devices {
		if device.Brand == brand {
			result = append(result, device)
		}
	}
	return result
}
