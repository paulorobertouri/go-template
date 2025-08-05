package user

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockRepository is a mock implementation of Repository for testing.
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByID(id int) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockRepository) Create(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) Update(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestService_GetUser(t *testing.T) {
	t.Run("valid user ID", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		expectedUser := &User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}

		mockRepo.On("GetByID", 1).Return(expectedUser, nil)

		result, err := service.GetUser(1)

		require.NoError(t, err)
		assert.Equal(t, expectedUser, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		_, err := service.GetUser(0)

		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrInvalidUserID)
		mockRepo.AssertNotCalled(t, "GetByID")
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		mockRepo.On("GetByID", 1).Return(nil, errors.New("database error"))

		_, err := service.GetUser(1)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get user")
		mockRepo.AssertExpectations(t)
	})
}

func TestService_CreateUser(t *testing.T) {
	t.Run("valid user data", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(nil)

		result, err := service.CreateUser("John Doe", "john@example.com")

		require.NoError(t, err)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "john@example.com", result.Email)
		assert.False(t, result.CreatedAt.IsZero())
		assert.False(t, result.UpdatedAt.IsZero())
		mockRepo.AssertExpectations(t)
	})

	t.Run("empty name", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		_, err := service.CreateUser("", "john@example.com")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("empty email", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		_, err := service.CreateUser("John Doe", "")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "email is required")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("invalid email", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		_, err := service.CreateUser("John Doe", "invalid-email")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid email format")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(errors.New("database error"))

		_, err := service.CreateUser("John Doe", "john@example.com")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create user")
		mockRepo.AssertExpectations(t)
	})
}

func TestService_UpdateUser(t *testing.T) {
	t.Run("valid update", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		existingUser := &User{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now().Add(-time.Hour),
			UpdatedAt: time.Now().Add(-time.Hour),
		}

		mockRepo.On("GetByID", 1).Return(existingUser, nil)
		mockRepo.On("Update", mock.AnythingOfType("*user.User")).Return(nil)

		result, err := service.UpdateUser(1, "Jane Doe", "jane@example.com")

		require.NoError(t, err)
		assert.Equal(t, "Jane Doe", result.Name)
		assert.Equal(t, "jane@example.com", result.Email)
		assert.True(t, result.UpdatedAt.After(existingUser.UpdatedAt))
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		_, err := service.UpdateUser(0, "Jane Doe", "jane@example.com")

		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrInvalidUserID)
		mockRepo.AssertNotCalled(t, "GetByID")
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		mockRepo.On("GetByID", 1).Return(nil, errors.New(ErrUserNotFound))

		_, err := service.UpdateUser(1, "Jane Doe", "jane@example.com")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get user")
		mockRepo.AssertExpectations(t)
	})
}

func TestService_DeleteUser(t *testing.T) {
	t.Run("valid deletion", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		mockRepo.On("Delete", 1).Return(nil)

		err := service.DeleteUser(1)

		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		mockRepo := new(MockRepository)
		service := NewService(mockRepo)

		err := service.DeleteUser(0)

		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrInvalidUserID)
		mockRepo.AssertNotCalled(t, "Delete")
	})
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "user@example.com", true},
		{"valid email with subdomain", "user@mail.example.com", true},
		{"invalid - no @", "userexample.com", false},
		{"invalid - no domain", "user@", false},
		{"invalid - no user", "@example.com", false},
		{"invalid - no dot after @", "user@example", false},
		{"invalid - too short", "a@", false},
		{"invalid - empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidEmail(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInMemoryRepository(t *testing.T) {
	t.Run("full CRUD operations", func(t *testing.T) {
		repo := NewInMemoryRepository()

		// Test Create
		user := &User{
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.Create(user)
		require.NoError(t, err)
		assert.Equal(t, 1, user.ID)

		// Test GetByID
		retrieved, err := repo.GetByID(1)
		require.NoError(t, err)
		assert.Equal(t, user, retrieved)

		// Test Update
		user.Name = "Jane Doe"
		err = repo.Update(user)
		require.NoError(t, err)

		updated, err := repo.GetByID(1)
		require.NoError(t, err)
		assert.Equal(t, "Jane Doe", updated.Name)

		// Test Delete
		err = repo.Delete(1)
		require.NoError(t, err)

		_, err = repo.GetByID(1)
		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrUserNotFound)
	})

	t.Run("get non-existent user", func(t *testing.T) {
		repo := NewInMemoryRepository()

		_, err := repo.GetByID(999)
		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrUserNotFound)
	})

	t.Run("update non-existent user", func(t *testing.T) {
		repo := NewInMemoryRepository()

		user := &User{ID: 999, Name: "Test"}
		err := repo.Update(user)
		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrUserNotFound)
	})

	t.Run("delete non-existent user", func(t *testing.T) {
		repo := NewInMemoryRepository()

		err := repo.Delete(999)
		require.Error(t, err)
		assert.Contains(t, err.Error(), ErrUserNotFound)
	})
}
