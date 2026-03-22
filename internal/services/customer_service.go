package services

import (
	"github.com/example/go-template/internal/domain"
	"github.com/example/go-template/internal/repositories"
)

// CustomerService handles customer business logic
type CustomerService struct {
	repo repositories.CustomerRepository
}

// NewCustomerService creates a new customer service
func NewCustomerService(repo repositories.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

// GetCustomer gets a customer by ID
func (s *CustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.GetCustomer(id)
}

// ListCustomers lists all customers
func (s *CustomerService) ListCustomers() ([]*domain.Customer, error) {
	return s.repo.ListCustomers()
}
