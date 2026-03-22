package unit

import (
	"testing"

	"github.com/example/go-template/internal/services"
)

const testUserSubject = "test-user"

func TestIssueToken(t *testing.T) {
	authService := services.NewAuthService()
	token, err := authService.IssueToken(testUserSubject)

	if err != nil {
		t.Fatalf("IssueToken failed: %v", err)
	}

	if token == "" {
		t.Error("Expected non-empty token")
	}
}

func TestValidateToken(t *testing.T) {
	authService := services.NewAuthService()
	token, _ := authService.IssueToken(testUserSubject)

	claims, err := authService.ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.Sub != testUserSubject {
		t.Errorf("Expected sub to be '%s', got %s", testUserSubject, claims.Sub)
	}
}

func TestValidateInvalidToken(t *testing.T) {
	authService := services.NewAuthService()
	_, err := authService.ValidateToken("invalid.token.here")

	if err == nil {
		t.Error("Expected validation to fail for invalid token")
	}
}

func TestValidateTokenWithWrongSecret(t *testing.T) {
	t.Setenv("JWT_SECRET", "your-super-secret-jwt-key-at-least-32-characters-long-for-hs256")
	authService := services.NewAuthService()
	token, _ := authService.IssueToken(testUserSubject)

	t.Setenv("JWT_SECRET", "different-secret-key-at-least-32-characters-long-for-hs256")
	authServiceWithDifferentSecret := services.NewAuthService()

	_, err := authServiceWithDifferentSecret.ValidateToken(token)
	if err == nil {
		t.Error("Expected validation to fail with different secret")
	}
}
