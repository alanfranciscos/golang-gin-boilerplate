package web

import (
	"github.com/alanf/go-boilerplate/internal/application"
	"github.com/gin-gonic/gin"
	"github.com/alanfranciscos/otel-lgtm-sdk-go/pkg/telemetry/middleware"
)

type Server struct {
	router        *gin.Engine
	port          string
	appName       string
	healthService application.HealthService
}

func NewServer(port string, appName string, healthService application.HealthService) *Server {
	s := &Server{
		router:        gin.Default(),
		port:          port,
		appName:       appName,
		healthService: healthService,
	}
	s.setupMiddlewares()
	s.setupRoutes()
	return s
}

func (s *Server) setupMiddlewares() {
	s.router.Use(middleware.GinMiddleware(s.appName))
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", s.handleHealth)
}

func (s *Server) handleHealth(c *gin.Context) {
	health := s.healthService.GetHealth(c.Request.Context())
	c.JSON(200, health)
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.port)
}
