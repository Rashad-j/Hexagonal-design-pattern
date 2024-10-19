package domain

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateDeviceInput struct {
	Name  string `json:"name" binding:"required"`
	Brand string `json:"brand" binding:"required"`
}
