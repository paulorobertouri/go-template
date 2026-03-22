package services

import (
	"testing"

	"github.com/example/go-template/internal/domain"
)

func TestIssueToken(t *testing.T) {
	authService := NewAuthService()
	token, err := authService.IssueToken("test-user")

	if err != nil {
		t.Fatalf("IssueToken failed: %v", err)
	}

	if token == "" {
		t.Error("Expected non-empty token")
	}
}

func TestValidateToken(t *testing.T) {
	authService := NewAuthService()
	token, _ := authService.IssueToken("test-user")

	claims, err := authService.ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.Sub != "test-user" {
		t.Errorf("Expected sub to be 'test-user', got %s", claims.Sub)
	}
}

func TestValidateInvalidToken(t *testing.T) {
	authService := NewAuthService()
	_, err := authService.ValidateToken("invalid.token.here")

	if err == nil {
		t.Error("Expected validation to fail for invalid token")
	}
}

func TestValidateTokenWithWrongSecret(t *testing.T) {
	authService := NewAuthService()
	token, _ := authService.IssueToken("test-user")

	// Create a new service with different secret
	authServiceWithDifferentSecret := &AuthService{
		secret:    "different-secret-key-at-least-32-characters-long-for-hs256",
		algorithm: "HS256",
		expiresIn: 3600,
	}

	_, err := authServiceWithDifferentSecret.ValidateToken(token)
	if err == nil {
		t.Error("Expected validation to fail with different secret")
	}
}
