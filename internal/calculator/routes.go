package calculator

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulorobertouri/go-template/internal/common"
)

// Response represents a calculation response
type Response struct {
	Result float64 `json:"result"`
}

// Handler handles calculator HTTP requests
type Handler struct {
	common.BaseHandler
	calc *Calculator
}

// NewHandler creates a new calculator handler
func NewHandler() *Handler {
	return &Handler{
		calc: New(),
	}
}

// RegisterRoutes registers calculator routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	routes := common.RouteGroup{
		Prefix: "",
		Routes: []common.Route{
			common.NamedRoute("/add/{a}/{b}", "GET", "calculator.add", h.handleAdd),
			common.NamedRoute("/subtract/{a}/{b}", "GET", "calculator.subtract", h.handleSubtract),
			common.NamedRoute("/multiply/{a}/{b}", "GET", "calculator.multiply", h.handleMultiply),
			common.NamedRoute("/divide/{a}/{b}", "GET", "calculator.divide", h.handleDivide),
		},
	}
	common.RegisterGroup(router, routes)
}

// handleAdd handles addition
// @Summary Add two numbers
// @Description Performs addition of two floating-point numbers
// @Tags calculator
// @Accept json
// @Produce json
// @Param a path number true "First number"
// @Param b path number true "Second number"
// @Success 200 {object} Response
// @Failure 400 {object} common.ErrorResponse
// @Router /add/{a}/{b} [get]
func (h *Handler) handleAdd(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.WriteBadRequest(w, err.Error())
		return
	}
	result := h.calc.Add(a, b)
	h.WriteSuccess(w, Response{Result: result})
}

// handleSubtract handles subtraction
// @Summary Subtract two numbers
// @Description Performs subtraction of two floating-point numbers (a - b)
// @Tags calculator
// @Accept json
// @Produce json
// @Param a path number true "First number (minuend)"
// @Param b path number true "Second number (subtrahend)"
// @Success 200 {object} Response
// @Failure 400 {object} common.ErrorResponse
// @Router /subtract/{a}/{b} [get]
func (h *Handler) handleSubtract(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.WriteBadRequest(w, err.Error())
		return
	}
	result := h.calc.Subtract(a, b)
	h.WriteSuccess(w, Response{Result: result})
}

// handleMultiply handles multiplication
// @Summary Multiply two numbers
// @Description Performs multiplication of two floating-point numbers
// @Tags calculator
// @Accept json
// @Produce json
// @Param a path number true "First number"
// @Param b path number true "Second number"
// @Success 200 {object} Response
// @Failure 400 {object} common.ErrorResponse
// @Router /multiply/{a}/{b} [get]
func (h *Handler) handleMultiply(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.WriteBadRequest(w, err.Error())
		return
	}
	result := h.calc.Multiply(a, b)
	h.WriteSuccess(w, Response{Result: result})
}

// handleDivide handles division
// @Summary Divide two numbers
// @Description Performs division of two floating-point numbers (a / b)
// @Tags calculator
// @Accept json
// @Produce json
// @Param a path number true "Dividend"
// @Param b path number true "Divisor (cannot be zero)"
// @Success 200 {object} Response
// @Failure 400 {object} common.ErrorResponse
// @Router /divide/{a}/{b} [get]
func (h *Handler) handleDivide(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.WriteBadRequest(w, err.Error())
		return
	}
	result, err := h.calc.Divide(a, b)
	if err != nil {
		h.WriteBadRequest(w, err.Error())
		return
	}
	h.WriteSuccess(w, Response{Result: result})
}

// getNumbers extracts numbers from URL parameters
func (h *Handler) getNumbers(r *http.Request) (float64, float64, error) {
	params := h.GetURLParams(r)

	a, err := params.Float64("a")
	if err != nil {
		return 0, 0, err
	}

	b, err := params.Float64("b")
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}
