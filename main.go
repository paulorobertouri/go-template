package main

import (
	"log"
	"os"

	"github.com/example/go-template/internal/api"
	"github.com/example/go-template/internal/di"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Initialize DI providers
	providers := di.NewProviders()

	// Register routes
	api.RegisterRoutes(e, providers)

	// Start server
	log.Printf("Starting server on port %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}
