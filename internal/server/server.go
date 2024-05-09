package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"orders/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Routable interface {
	RegisterRoutes(router *gin.Engine)
}

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	port := config.Env("PORT")
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return &Server{
		router:     router,
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	log.Printf("Server started on %s", s.httpServer.Addr)

	<-ctx.Done()

	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Server has been shutdown gracefully.")

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Server) RegisterRoutes(apis []Routable) {
	for _, api := range apis {
		api.RegisterRoutes(s.router)
	}
}
