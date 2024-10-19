
# Mobile Device Management API

This is a RESTful API built in Go using the Gin framework. It supports basic mobile device management operations such as adding, updating, deleting, and listing devices in store. The architecture follows a Hexagonal (Ports and Adapters) design pattern, making the system scalable and maintainable.

## Table of Contents

- [Mobile Device Management API](#mobile-device-management-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Project Structure](#project-structure)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
    - [Run Locally](#run-locally)
    - [Run with Docker](#run-with-docker)
  - [API Endpoints](#api-endpoints)
    - [Example Request Body for Adding a Device](#example-request-body-for-adding-a-device)
  - [Testing](#testing)
  - [OpenAPI Documentation](#openapi-documentation)
    - [Serve Swagger UI](#serve-swagger-ui)
      - [Steps to Generate Swagger Docs:](#steps-to-generate-swagger-docs)
  - [Future Improvements](#future-improvements)
  - [License](#license)
  - [Author](#author)

## Features

- Add a new device
- Get a device by its ID
- List all devices
- Update a device (full or partial update)
- Delete a device
- Search for devices by brand
- Follows the Hexagonal Architecture pattern (Ports and Adapters)
- Includes OpenAPI (Swagger) documentation

## Project Structure

The project is utilizing a hexagonal architecture (Ports and Adapters) to ensure flexibility and easy extension. Below is a breakdown of the structure:

```
device-management-api/
├── cmd/
│   └── main.go                  # Entry point for the application
├── internal/
│   ├── core/
│   │   ├── domain/
│   │   │   └── device.go         # Domain model
│   │   └── ports/
│   │       └── device_repo.go    # Repository interface (Port)
│   ├── adapters/
│   │   ├── repository/
│   │   │   └── memory_repo.go    # In-memory repository (Adapter)
│   │   └── http/
│   │       ├── device_handler.go # HTTP handler for devices (Adapter)
│   │       └── server.go         # Generic server logic
│   └── usecases/
│       └── device_service.go     # Business logic (Use case)
├── tests/
│   └── device_service_test.go    # Unit tests for services
├── Dockerfile                    # Docker configuration
├── go.mod                         # Go modules
├── go.sum
└── README.md                     # Documentation
```

## Prerequisites

- Go 1.22+
- Docker (optional for containerized setup)
- [Swagger UI](https://swagger.io/tools/swagger-ui/) for API documentation

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/rashad-j/device-management-api.git
   cd device-management-api
   ```

2. **Install dependencies:**

   Ensure that Go modules are enabled and install dependencies using:

   ```bash
   go mod tidy
   ```

## Running the Application

### Run Locally

You can run the API locally by executing the following command:

```bash
go run cmd/main.go
```

The API will be available at [http://localhost:8080](http://localhost:8080).

### Run with Docker

To run the API in a Docker container:

1. **Build the Docker image:**

   ```bash
   make docker-build
   ```

2. **Run the Docker container:**

   ```bash
   make docker-run
   ```

Now the API should be running at [http://localhost:8080/v1](http://localhost:8080).

## API Endpoints

The API provides the following endpoints:

| HTTP Method | Endpoint            | Description                      |
|-------------|---------------------|----------------------------------|
| `POST`      | `/v1/devices`           | Add a new device                 |
| `GET`       | `/v1/devices`           | List all devices                 |
| `GET`       | `/v1/devices/{id}`      | Get a device by its ID           |
| `PUT`       | `/v1/devices/{id}`      | Update a device by its ID        |
| `DELETE`    | `/v1/devices/{id}`      | Delete a device by its ID        |
| `GET`       | `/v1/devices/search`    | Search for devices by brand      |

### Example Request Body for Adding a Device

```json
{
  "name": "iPhone 12",
  "brand": "Apple"
}
```

## Testing

The project includes unit tests for the business logic. You can run the tests using the following command:

```bash
go test ./...
```

Make sure to maintain good test coverage as you add new features to the API.


## OpenAPI Documentation

The API is documented using OpenAPI (formerly Swagger). You can view the documentation by visiting the `/swagger` endpoint in the browser.

### Serve Swagger UI

To serve the Swagger UI, navigate to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) after starting the application.

Swagger documentation is generated automatically using the **Swaggo** library.

#### Steps to Generate Swagger Docs:

1. **Install Swaggo:**

   ```bash
   go get -u github.com/swaggo/swag/cmd/swag
   ```

2. **Generate Swagger docs:**

   ```bash
   swag init
   ```

The generated Swagger specs will be placed in the `docs` folder, and they can be served automatically through the `/swagger` endpoint.

## Future Improvements

- Add more use cases (e.g., User Management) to test the flexibility of the system.
- Add support for persistent storage (e.g., PostgreSQL, MongoDB) instead of the in-memory repository.
- Implement authentication and authorization for better security.
- Improve unit test coverage and add integration tests.
- Add logging and monitoring tools.
  
---

## License

This project is open-source and licensed under the [MIT License](LICENSE).

---

## Author

- **Your Name**
- GitHub: [rashad-j](https://github.com/rashad-j)

Feel free to fork this repository and contribute by opening a pull request!

---

