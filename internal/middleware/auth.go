package middleware

import (
	"net/http"
	"strings"

	"github.com/example/go-template/internal/services"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware validates JWT tokens from Authorization header
func JWTMiddleware(authService *services.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing authorization header",
				})
			}

			// Extract Bearer token
			const bearerPrefix = "Bearer "
			if !strings.HasPrefix(authHeader, bearerPrefix) {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid authorization header format",
				})
			}

			token := strings.TrimPrefix(authHeader, bearerPrefix)

			// Validate token
			claims, err := authService.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": err.Error(),
				})
			}

			// Store claims in context for later use
			c.Set("user", claims.Sub)
			return next(c)
		}
	}
}
