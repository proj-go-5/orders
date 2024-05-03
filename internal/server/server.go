package server

import (
	"github.com/gin-gonic/gin"
	"orders/internal/config"
)

type Routable interface {
	RegisterRoutes(router *gin.Engine)
}

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{router}
}

func (s *Server) Start() error {
	port := config.Env("PORT")
	err := s.router.Run(":" + port)

	return err
}

func (s *Server) RegisterRoutes(apis []Routable) {
	for _, api := range apis {
		api.RegisterRoutes(s.router)
	}
}
