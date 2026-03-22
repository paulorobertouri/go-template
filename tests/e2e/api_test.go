package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/example/go-template/internal/api"
	"github.com/example/go-template/internal/di"
	"github.com/example/go-template/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupTestServer() *echo.Echo {
	e := echo.New()
	providers := di.NewProviders()
	api.RegisterRoutes(e, providers)
	return e
}

func TestPublicEndpoint(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/v1/public", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response domain.PublicResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "This is a public endpoint", response.Message)
}

func TestLoginEndpoint(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response domain.LoginResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEmpty(t, response.Token)

	// Check headers
	assert.NotEmpty(t, rec.Header().Get("X-JWT-Token"))
	assert.NotEmpty(t, rec.Header().Get("Authorization"))
}

func TestPrivateEndpointWithoutAuth(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/v1/private", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestPrivateEndpointWithAuth(t *testing.T) {
	e := setupTestServer()
	providers := di.NewProviders()

	// Get token
	token, _ := providers.AuthService.IssueToken("test-user")

	// Call private endpoint with token
	req := httptest.NewRequest(http.MethodGet, "/v1/private", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response domain.PrivateResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "This is a private endpoint", response.Message)
	assert.Equal(t, "test-user", response.User)
}

func TestCustomerEndpoint(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/v1/customer", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response []domain.CustomerResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEmpty(t, response)
}
