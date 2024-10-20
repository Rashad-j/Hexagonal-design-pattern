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
	tests := []struct {
		name           string
		body           string
		expectedName   string
		expectedBrand  string
		expectedStatus int
	}{
		{
			name:           "valid device creation",
			body:           `{"name":"iPhone 12","brand":"Apple"}`,
			expectedName:   "iPhone 12",
			expectedBrand:  "Apple",
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockDeviceRepository(ctrl)
			service := usecases.NewDeviceService(mockRepo)
			handler := internalHttp.NewDeviceHandler(service)

			go func() {
				server := internalHttp.NewServer(config.LoadConfig())
				server.RegisterHandler(handler)
				server.Start()
			}()

			// Wait for the server to start
			for range [5]struct{}{} {
				_, err := http.Get("http://localhost:8080/v1/ping")
				if err == nil {
					break
				}
				fmt.Println("Waiting for server to start...")
				time.Sleep(1 * time.Second)
			}

			mockRepo.EXPECT().Create(gomock.Any()).Times(1)

			payload := io.NopCloser(bytes.NewBufferString(tt.body))
			req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/v1/devices", payload)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var device domain.Device
			err = json.NewDecoder(resp.Body).Decode(&device)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedName, device.Name)
			assert.Equal(t, tt.expectedBrand, device.Brand)
		})
	}
}

func TestDeviceHandler_GetDeviceByID(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		deviceID       uuid.UUID
		expectedDevice domain.Device
		expectedStatus int
	}{
		{
			name:     "valid device retrieval",
			deviceID: uuid.New(),
			expectedDevice: domain.Device{
				Name:  "iPhone 12",
				Brand: "Apple",
			},
			expectedStatus: http.StatusOK,
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockDeviceRepository(ctrl)
			service := usecases.NewDeviceService(mockRepo)
			handler := internalHttp.NewDeviceHandler(service)

			go func() {
				cfg := config.LoadConfig().WithPort("8081")
				server := internalHttp.NewServer(cfg)
				server.RegisterHandler(handler)
				server.Start()
			}()

			// Wait for the server to start
			for range [5]struct{}{} {
				_, err := http.Get("http://localhost:8081/v1/ping")
				if err == nil {
					break
				}
				fmt.Println("Waiting for server to start...")
				time.Sleep(1 * time.Second)
			}

			mockRepo.EXPECT().GetById(tt.deviceID).Return(tt.expectedDevice, nil).Times(1)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8081/v1/devices/%s", tt.deviceID), nil)
			if err != nil {
				t.Fatal(err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var device domain.Device
			err = json.NewDecoder(resp.Body).Decode(&device)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedDevice.Name, device.Name)
			assert.Equal(t, tt.expectedDevice.Brand, device.Brand)
		})
	}
}
