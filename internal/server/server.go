package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulorobertouri/go-template/internal/calculator"
	"github.com/paulorobertouri/go-template/internal/common"
	"github.com/paulorobertouri/go-template/internal/greeting"
	"github.com/paulorobertouri/go-template/internal/user"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Server represents the HTTP server
type Server struct {
	common.BaseHandler
	router            *mux.Router
	calculatorHandler *calculator.Handler
	userHandler       *user.Handler
	greetingHandler   *greeting.Handler
}

// New creates a new server instance
func New() *Server {
	// Create user service
	userService := user.NewService()

	s := &Server{
		router:            mux.NewRouter(),
		calculatorHandler: calculator.NewHandler(),
		userHandler:       user.NewHandler(userService),
		greetingHandler:   greeting.NewHandler(),
	}

	s.setupRoutes()
	return s
}

// Router returns the server's router
func (s *Server) Router() http.Handler {
	return s.router
}

// setupRoutes configures the server routes
func (s *Server) setupRoutes() {
	// Health and root routes
	s.router.HandleFunc("/health", s.handleHealth).Methods("GET")
	s.router.HandleFunc("/", s.handleRoot).Methods("GET")

	// Swagger UI
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Register all handlers
	s.calculatorHandler.RegisterRoutes(s.router)
	s.userHandler.RegisterRoutes(s.router)
	s.greetingHandler.RegisterRoutes(s.router)
}

// handleHealth handles health check requests
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	s.WriteSuccess(w, map[string]string{"status": "ok"})
}

// handleRoot handles root requests
// @Summary Welcome endpoint
// @Description Returns a welcome message for the API
// @Tags general
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	s.WriteSuccess(w, map[string]string{"message": "Welcome to the Go Template API"})
}
