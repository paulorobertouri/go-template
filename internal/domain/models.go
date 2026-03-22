package domain

// Customer represents a customer in the system
type Customer struct {
	ID    string
	Name  string
	Email string
}

// LoginRequest represents a login request
type LoginRequest struct{}

// LoginResponse represents a login response with JWT token
type LoginResponse struct {
	Token string `json:"token"`
}

// PrivateResponse represents a protected endpoint response
type PrivateResponse struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

// PublicResponse represents a public endpoint response
type PublicResponse struct {
	Message string `json:"message"`
}

// CustomerResponse represents a customer response
type CustomerResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	Sub string    `json:"sub"`
	Iat int64     `json:"iat"`
	Exp int64     `json:"exp"`
}
