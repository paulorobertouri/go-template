// Package user provides simple user management functionality.
package user

import (
	"errors"
)

const (
	ErrInvalidUserID = "invalid user ID"
	ErrUserNotFound  = "user not found"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Service provides user operations
type Service struct {
	users  map[int]*User
	nextID int
}

// NewService creates a new user service
func NewService() *Service {
	return &Service{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New(ErrInvalidUserID)
	}

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New(ErrUserNotFound)
	}

	return user, nil
}

// CreateUser creates a new user
func (s *Service) CreateUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if email == "" {
		return nil, errors.New("email is required")
	}

	user := &User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
	}

	s.users[s.nextID] = user
	s.nextID++

	return user, nil
}

// UpdateUser updates an existing user
func (s *Service) UpdateUser(id int, name, email string) (*User, error) {
	if id <= 0 {
		return nil, errors.New(ErrInvalidUserID)
	}

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New(ErrUserNotFound)
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	if email == "" {
		return nil, errors.New("email is required")
	}

	user.Name = name
	user.Email = email

	return user, nil
}

// DeleteUser deletes a user by ID
func (s *Service) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New(ErrInvalidUserID)
	}

	if _, exists := s.users[id]; !exists {
		return errors.New(ErrUserNotFound)
	}

	delete(s.users, id)
	return nil
}

// GetAllUsers returns all users
func (s *Service) GetAllUsers() []*User {
	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}
