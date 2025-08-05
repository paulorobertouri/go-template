package greeting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandlerHello(t *testing.T) {
	tests := []struct {
		name           string
		urlPath        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid name",
			urlPath:        "/greeting/John",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":{"message":"Hello, John!"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create handler
			handler := NewHandler()

			// Create router and register routes
			router := mux.NewRouter()
			handler.RegisterRoutes(router)

			// Create request
			req := httptest.NewRequest("GET", tt.urlPath, nil)
			w := httptest.NewRecorder()

			// Serve HTTP
			router.ServeHTTP(w, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestHandlerFormalGreeting(t *testing.T) {
	tests := []struct {
		name           string
		urlPath        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid formal greeting",
			urlPath:        "/greeting/formal/Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":{"message":"Good day, Alice. It's a pleasure to meet you."}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create handler
			handler := NewHandler()

			// Create router and register routes
			router := mux.NewRouter()
			handler.RegisterRoutes(router)

			// Create request
			req := httptest.NewRequest("GET", tt.urlPath, nil)
			w := httptest.NewRecorder()

			// Serve HTTP
			router.ServeHTTP(w, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
