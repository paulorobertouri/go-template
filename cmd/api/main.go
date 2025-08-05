package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/paulorobertouri/go-template/docs" // Import docs for swagger
	"github.com/paulorobertouri/go-template/internal/server"
)

// @title Go Template API
// @version 1.0
// @description A simple Go API template with calculator and user management endpoints
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Create a new server instance
	srv := server.New()

	// Create HTTP server
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv.Router(),
	}

	// Channel to listen for interrupt signal to terminate server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		log.Println("Swagger UI available at: http://localhost:8080/swagger/")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
