// tests/device_handler_test.go

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	internalHttp "github.com/rashad-j/device-management-api/internal/adapters/http"
	"github.com/rashad-j/device-management-api/internal/config"
	"github.com/rashad-j/device-management-api/internal/core/domain"
	"github.com/rashad-j/device-management-api/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func TestDeviceHandler_CreateDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)
	service := usecases.NewDeviceService(mockRepo)
	handler := internalHttp.NewDeviceHandler(service)
	go func() {
		server := internalHttp.NewServer(config.LoadConfig())
		server.RegisterHandler(handler)

		// start server
		server.Start()
	}()

	// keep retrying until the server is up
	for range [5]struct{}{} {
		_, err := http.Get("http://localhost:8080/v1/ping")
		if err == nil {
			break
		}
		fmt.Println("Waiting for server to start...")
		time.Sleep(1 * time.Second)
	}

	mockRepo.EXPECT().Create(gomock.Any()).Times(1)

	body := `{"name":"iPhone 12","brand":"Apple"}`
	payload := io.NopCloser(bytes.NewBufferString(body))
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/v1/devices", payload)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// execute http request to local server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	var device domain.Device
	err = json.NewDecoder(resp.Body).Decode(&device)
	assert.Nil(t, err)
	assert.Equal(t, "iPhone 12", device.Name)
	assert.Equal(t, "Apple", device.Brand)

}

func TestDeviceHandler_GetDeviceByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDeviceRepository(ctrl)
	service := usecases.NewDeviceService(mockRepo)
	handler := internalHttp.NewDeviceHandler(service)

	go func() {
		server := internalHttp.NewServer(config.LoadConfig())
		server.RegisterHandler(handler)

		// start server
		server.Start()
	}()

	// keep retrying until the server is up
	for range [5]struct{}{} {
		_, err := http.Get("http://localhost:8080/v1/ping")
		if err == nil {
			break
		}
		fmt.Println("Waiting for server to start...")
		time.Sleep(1 * time.Second)
	}

	id := uuid.New()
	expectedDevice := domain.Device{
		ID:    id,
		Name:  "iPhone 12",
		Brand: "Apple",
	}

	mockRepo.EXPECT().GetById(id).Return(expectedDevice, nil).Times(1)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v1/devices/%s", id), nil)
	if err != nil {
		t.Fatal(err)
	}
	// execute http request to local server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Assertions
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var device domain.Device
	err = json.NewDecoder(resp.Body).Decode(&device)
	assert.Nil(t, err)
	assert.Equal(t, "iPhone 12", device.Name)
	assert.Equal(t, "Apple", device.Brand)
}
