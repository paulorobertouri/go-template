package greeting

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulorobertouri/go-template/internal/common"
)

// GreetingResponse represents a greeting response
type GreetingResponse struct {
	Message string `json:"message"`
}

// Handler handles greeting HTTP requests
type Handler struct {
	common.BaseHandler
}

// NewHandler creates a new greeting handler
func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes registers greeting routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	routes := common.RouteGroup{
		Prefix: "/greeting",
		Routes: []common.Route{
			common.NamedRoute("/{name}", "GET", "greeting.hello", h.handleHello),
			common.NamedRoute("/formal/{name}", "GET", "greeting.formal", h.handleFormalGreeting),
		},
	}
	common.RegisterGroup(router, routes)
}

// handleHello handles simple greeting
// @Summary Greet a person
// @Description Returns a simple greeting message for the given name
// @Tags greeting
// @Accept json
// @Produce json
// @Param name path string true "Person's name"
// @Success 200 {object} GreetingResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /greeting/{name} [get]
func (h *Handler) handleHello(w http.ResponseWriter, r *http.Request) {
	params := h.GetURLParams(r)
	name, exists := params.String("name")
	if !exists || name == "" {
		h.WriteBadRequest(w, "name is required")
		return
	}

	response := GreetingResponse{
		Message: "Hello, " + name + "!",
	}
	h.WriteSuccess(w, response)
}

// handleFormalGreeting handles formal greeting
// @Summary Formal greeting
// @Description Returns a formal greeting message for the given name
// @Tags greeting
// @Accept json
// @Produce json
// @Param name path string true "Person's name"
// @Success 200 {object} GreetingResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /greeting/formal/{name} [get]
func (h *Handler) handleFormalGreeting(w http.ResponseWriter, r *http.Request) {
	params := h.GetURLParams(r)
	name, exists := params.String("name")
	if !exists || name == "" {
		h.WriteBadRequest(w, "name is required")
		return
	}

	response := GreetingResponse{
		Message: "Good day, " + name + ". It's a pleasure to meet you.",
	}
	h.WriteSuccess(w, response)
}
