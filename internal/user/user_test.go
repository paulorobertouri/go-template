package user

import (
	"testing"
)

func TestService_GetUser(t *testing.T) {
	service := NewService()

	// Test getting a non-existent user
	_, err := service.GetUser(1)
	if err == nil {
		t.Error("Expected error for non-existent user")
	}

	// Create a user first
	user, err := service.CreateUser("John Doe", "john@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Test getting the created user
	retrievedUser, err := service.GetUser(user.ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if retrievedUser.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", retrievedUser.Name)
	}
}

func TestService_CreateUser(t *testing.T) {
	service := NewService()

	user, err := service.CreateUser("Jane Doe", "jane@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if user.Name != "Jane Doe" {
		t.Errorf("Expected name 'Jane Doe', got '%s'", user.Name)
	}

	if user.Email != "jane@example.com" {
		t.Errorf("Expected email 'jane@example.com', got '%s'", user.Email)
	}

	// Test creating user with empty name
	_, err = service.CreateUser("", "test@example.com")
	if err == nil {
		t.Error("Expected error for empty name")
	}
}

func TestService_UpdateUser(t *testing.T) {
	service := NewService()

	// Create a user first
	user, err := service.CreateUser("John Doe", "john@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Update the user
	updatedUser, err := service.UpdateUser(user.ID, "John Smith", "johnsmith@example.com")
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	if updatedUser.Name != "John Smith" {
		t.Errorf("Expected name 'John Smith', got '%s'", updatedUser.Name)
	}

	if updatedUser.Email != "johnsmith@example.com" {
		t.Errorf("Expected email 'johnsmith@example.com', got '%s'", updatedUser.Email)
	}
}

func TestService_DeleteUser(t *testing.T) {
	service := NewService()

	// Create a user first
	user, err := service.CreateUser("John Doe", "john@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Delete the user
	err = service.DeleteUser(user.ID)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Try to get the deleted user
	_, err = service.GetUser(user.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user")
	}
}
