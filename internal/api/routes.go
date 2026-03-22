package api

import (
	"net/http"

	"github.com/example/go-template/internal/di"
	"github.com/example/go-template/internal/domain"
	"github.com/example/go-template/internal/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, providers *di.Providers) {
	// Swagger documentation
	e.GET("/docs", echoSwagger.WrapHandler)
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// Public routes
	e.GET("/v1/public", handlePublic)
	e.GET("/v1/customer", handleGetCustomer(providers))

	// Auth routes
	e.POST("/v1/auth/login", handleLogin(providers))

	// Protected routes (with JWT middleware)
	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware(providers.AuthService))
	protected.GET("/v1/private", handlePrivate)
}

// handlePublic handles public endpoint
func handlePublic(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.PublicResponse{
		Message: "This is a public endpoint",
	})
}

// handleGetCustomer handles getting a customer
func handleGetCustomer(providers *di.Providers) echo.HandlerFunc {
	return func(c echo.Context) error {
		customers, err := providers.CustomerService.ListCustomers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		result := make([]domain.CustomerResponse, len(customers))
		for i, customer := range customers {
			result[i] = domain.CustomerResponse{
				ID:    customer.ID,
				Name:  customer.Name,
				Email: customer.Email,
			}
		}

		return c.JSON(http.StatusOK, result)
	}
}

// handleLogin handles user login and JWT token generation
func handleLogin(providers *di.Providers) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Generate JWT token
		token, err := providers.AuthService.IssueToken("user@example.com")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to generate token",
			})
		}

		// Set response headers
		c.Response().Header().Set("Authorization", "Bearer "+token)
		c.Response().Header().Set("X-JWT-Token", token)

		return c.JSON(http.StatusOK, domain.LoginResponse{
			Token: token,
		})
	}
}

// handlePrivate handles protected endpoint that requires JWT
func handlePrivate(c echo.Context) error {
	user := c.Get("user").(string)

	return c.JSON(http.StatusOK, domain.PrivateResponse{
		Message: "This is a private endpoint",
		User:    user,
	})
}
