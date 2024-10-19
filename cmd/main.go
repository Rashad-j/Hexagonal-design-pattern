// cmd/main.go
package main

import (
	"github.com/rashad-j/device-management-api/internal/adapters/http"
	"github.com/rashad-j/device-management-api/internal/adapters/repository"
	"github.com/rashad-j/device-management-api/internal/config"
	"github.com/rashad-j/device-management-api/internal/usecases"
)

func main() {
	// get env config
	cfg := config.LoadConfig()

	// init repository and service
	deviceRepo := repository.NewMemoryDeviceRepository()
	deviceService := usecases.NewDeviceService(deviceRepo)
	deviceHandler := http.NewDeviceHandler(deviceService)

	// create new http server
	server := http.NewServer(cfg)
	server.RegisterHandler(deviceHandler)

	// start server
	server.Start()
}
