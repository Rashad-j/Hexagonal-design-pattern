package http

import (
	"net/http"

	"github.com/rashad-j/device-management-api/internal/core/domain"
	"github.com/rashad-j/device-management-api/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeviceHandler struct {
	service *usecases.DeviceService
}

func NewDeviceHandler(service *usecases.DeviceService) *DeviceHandler {
	return &DeviceHandler{service: service}
}

func (h *DeviceHandler) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1.POST("/devices", h.CreateDevice)
	v1.GET("/devices/:id", h.GetDeviceById)
	v1.GET("/devices", h.ListDevices)
	v1.PUT("/devices/:id", h.UpdateDevice)
	v1.DELETE("/devices/:id", h.DeleteDevice)
	v1.GET("/devices/search", h.SearchDevicesByBrand)
}

func (h *DeviceHandler) CreateDevice(c *gin.Context) {
	var input domain.CreateDeviceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := h.service.CreateDevice(input)
	c.JSON(http.StatusCreated, device)
}

func (h *DeviceHandler) GetDeviceById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	device, err := h.service.GetDeviceById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, device)
}

func (h *DeviceHandler) ListDevices(c *gin.Context) {
	devices := h.service.ListDevices()
	c.JSON(http.StatusOK, devices)
}

func (h *DeviceHandler) UpdateDevice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input domain.CreateDeviceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device, err := h.service.UpdateDevice(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, device)
}

func (h *DeviceHandler) DeleteDevice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteDevice(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *DeviceHandler) SearchDevicesByBrand(c *gin.Context) {
	brand := c.Query("brand")
	devices := h.service.SearchDevicesByBrand(brand)
	c.JSON(http.StatusOK, devices)
}
