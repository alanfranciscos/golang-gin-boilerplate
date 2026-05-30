package web

import (
	"github.com/alanf/go-boilerplate/internal/application"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router        *gin.Engine
	port          string
	healthService application.HealthService
}

func NewServer(port string, healthService application.HealthService) *Server {
	s := &Server{
		router:        gin.Default(),
		port:          port,
		healthService: healthService,
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", s.handleHealth)
}

func (s *Server) handleHealth(c *gin.Context) {
	health := s.healthService.GetHealth()
	c.JSON(200, health)
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.port)
}
