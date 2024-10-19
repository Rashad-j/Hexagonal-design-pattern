package http

import (
	"fmt"

	"github.com/rashad-j/device-management-api/internal/config"

	"github.com/gin-gonic/gin"
)

// Handler is a generic interface for HTTP handlers.
type Handler interface {
	RegisterRoutes(router *gin.Engine)
}

type Server struct {
	router   *gin.Engine
	handlers []Handler
	port     string
}

// NewServer initializes a new server instance with a generic list of handlers.
func NewServer(cfg *config.Config) *Server {
	router := gin.Default()
	return &Server{
		router:   router,
		handlers: []Handler{},
		port:     cfg.Port,
	}
}

// RegisterHandler allows registering a new handler for the server.
func (s *Server) RegisterHandler(handler Handler) {
	s.handlers = append(s.handlers, handler)
}

// SetupRoutes calls the RegisterRoutes method on each registered handler.
func (s *Server) setupRoutes() {
	for _, handler := range s.handlers {
		handler.RegisterRoutes(s.router)
	}
}

// Start starts the HTTP server on the specified port.
func (s *Server) Start() error {
	s.setupRoutes()
	fmt.Printf("Server running on port %s\n", s.port)
	return s.router.Run(fmt.Sprintf(":%s", s.port))
}
