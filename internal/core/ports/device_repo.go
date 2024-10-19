package ports

import (
	"github.com/rashad-j/device-management-api/internal/core/domain"

	"github.com/google/uuid"
)

type DeviceRepository interface {
	Create(device domain.Device)
	GetById(id uuid.UUID) (domain.Device, error)
	List() []domain.Device
	Update(id uuid.UUID, device domain.Device) error
	Delete(id uuid.UUID) error
	SearchByBrand(brand string) []domain.Device
}
