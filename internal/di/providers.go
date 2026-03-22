package di

import (
	"github.com/example/go-template/internal/repositories"
	"github.com/example/go-template/internal/services"
)

// Providers holds all service providers (dependency injection container)
type Providers struct {
	AuthService     *services.AuthService
	CustomerService *services.CustomerService
	CustomerRepo    repositories.CustomerRepository
}

// NewProviders initializes all providers
func NewProviders() *Providers {
	// Initialize repositories
	customerRepo := repositories.NewInMemoryCustomerRepository()

	// Initialize services
	authService := services.NewAuthService()
	customerService := services.NewCustomerService(customerRepo)

	return &Providers{
		AuthService:     authService,
		CustomerService: customerService,
		CustomerRepo:    customerRepo,
	}
}
