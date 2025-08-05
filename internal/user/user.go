// Package user provides user-related functionality.
package user

import (
	"errors"
	"fmt"
	"time"
)

// Error constants
const (
	ErrInvalidUserID = "invalid user ID"
	ErrUserNotFound  = "user not found"
)

// User represents a user in the system.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Repository defines the interface for user data operations.
type Repository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

// Service provides user business logic.
type Service struct {
	repo Repository
}

// NewService creates a new user service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetUser retrieves a user by ID.
func (s *Service) GetUser(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New(ErrInvalidUserID)
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// CreateUser creates a new user.
func (s *Service) CreateUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	
	if email == "" {
		return nil, errors.New("email is required")
	}

	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	user := &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// UpdateUser updates an existing user.
func (s *Service) UpdateUser(id int, name, email string) (*User, error) {
	if id <= 0 {
		return nil, errors.New(ErrInvalidUserID)
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if name != "" {
		user.Name = name
	}
	
	if email != "" {
		if !isValidEmail(email) {
			return nil, errors.New("invalid email format")
		}
		user.Email = email
	}

	user.UpdatedAt = time.Now()

	if err := s.repo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// DeleteUser deletes a user by ID.
func (s *Service) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New(ErrInvalidUserID)
	}

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// isValidEmail performs basic email validation.
func isValidEmail(email string) bool {
	// Simple email validation for demonstration
	// In production, use a proper email validation library
	if len(email) < 3 {
		return false
	}
	
	atFound := false
	dotAfterAt := false
	
	for i, char := range email {
		if char == '@' {
			if atFound || i == 0 || i == len(email)-1 {
				return false
			}
			atFound = true
		} else if char == '.' && atFound && i > 0 && i < len(email)-1 {
			dotAfterAt = true
		}
	}
	
	return atFound && dotAfterAt
}

// InMemoryRepository is a simple in-memory implementation of Repository.
type InMemoryRepository struct {
	users  map[int]*User
	nextID int
}

// NewInMemoryRepository creates a new in-memory repository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

// GetByID retrieves a user by ID.
func (r *InMemoryRepository) GetByID(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New(ErrUserNotFound)
	}
	return user, nil
}

// Create creates a new user.
func (r *InMemoryRepository) Create(user *User) error {
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}

// Update updates an existing user.
func (r *InMemoryRepository) Update(user *User) error {
	if _, exists := r.users[user.ID]; !exists {
		return errors.New(ErrUserNotFound)
	}
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user by ID.
func (r *InMemoryRepository) Delete(id int) error {
	if _, exists := r.users[id]; !exists {
		return errors.New(ErrUserNotFound)
	}
	delete(r.users, id)
	return nil
}
