package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/example/go-template/internal/domain"
)

// AuthService handles JWT authentication
type AuthService struct {
	secret    string
	algorithm string
	expiresIn int64
}

// NewAuthService creates a new authentication service
func NewAuthService() *AuthService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-super-secret-jwt-key-at-least-32-characters-long-for-hs256"
	}

	algorithm := os.Getenv("JWT_ALGORITHM")
	if algorithm == "" {
		algorithm = "HS256"
	}

	expiresIn := int64(3600)
	if exp := os.Getenv("JWT_EXPIRATION"); exp != "" {
		if parsedExp, err := strconv.ParseInt(exp, 10, 64); err == nil {
			expiresIn = parsedExp
		}
	}

	return &AuthService{
		secret:    secret,
		algorithm: algorithm,
		expiresIn: expiresIn,
	}
}

// IssueToken generates a JWT token
func (s *AuthService) IssueToken(subject string) (string, error) {
	now := time.Now().Unix()
	exp := now + s.expiresIn

	claims := domain.TokenClaims{
		Sub: subject,
		Iat: now,
		Exp: exp,
	}

	// Create JWT header
	header := map[string]string{
		"alg": s.algorithm,
		"typ": "JWT",
	}

	headerJSON, _ := json.Marshal(header)
	headerB64 := base64.RawURLEncoding.EncodeToString(headerJSON)

	// Create JWT payload
	payloadJSON, _ := json.Marshal(claims)
	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadJSON)

	// Create signature
	message := headerB64 + "." + payloadB64
	signature := s.signHS256(message)

	token := message + "." + signature
	return token, nil
}

// ValidateToken validates a JWT token and returns claims
func (s *AuthService) ValidateToken(tokenString string) (*domain.TokenClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	headerB64 := parts[0]
	payloadB64 := parts[1]
	signatureB64 := parts[2]

	// Verify signature
	message := headerB64 + "." + payloadB64
	expectedSignature := s.signHS256(message)
	if signatureB64 != expectedSignature {
		return nil, fmt.Errorf("invalid token signature")
	}

	// Decode and parse claims
	payloadJSON, err := base64.RawURLEncoding.DecodeString(payloadB64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	var claims domain.TokenClaims
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	// Check expiration
	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token expired")
	}

	return &claims, nil
}

// signHS256 creates HMAC-SHA256 signature
func (s *AuthService) signHS256(message string) string {
	h := hmac.New(sha256.New, []byte(s.secret))
	h.Write([]byte(message))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return signature
}
