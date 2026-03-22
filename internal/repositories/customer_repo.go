package repositories

import (
	"github.com/example/go-template/internal/domain"
)

// CustomerRepository interface for customer data access
type CustomerRepository interface {
	GetCustomer(id string) (*domain.Customer, error)
	ListCustomers() ([]*domain.Customer, error)
}

// InMemoryCustomerRepository implements CustomerRepository with in-memory storage
type InMemoryCustomerRepository struct {
	customers map[string]*domain.Customer
}

// NewInMemoryCustomerRepository creates a new in-memory customer repository
func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: map[string]*domain.Customer{
			"1": {ID: "1", Name: "John Doe", Email: "john@example.com"},
			"2": {ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
		},
	}
}

// GetCustomer retrieves a customer by ID
func (r *InMemoryCustomerRepository) GetCustomer(id string) (*domain.Customer, error) {
	if customer, exists := r.customers[id]; exists {
		return customer, nil
	}
	return nil, nil
}

// ListCustomers returns all customers
func (r *InMemoryCustomerRepository) ListCustomers() ([]*domain.Customer, error) {
	customers := make([]*domain.Customer, 0, len(r.customers))
	for _, customer := range r.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}
