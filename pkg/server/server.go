package server

import (
	"fmt"

	"github.com/max71211/converter-service/pkg/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	s := Server{
		config: cfg,
		router: gin.Default(),
	}

	s.router.GET("/", s.handleIndex)
	s.router.GET("/metrics", s.handleMetrics)

	return &s
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))
}
