package calculator

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Request represents a calculation request
type Request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

// Response represents a calculation response
type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Handler handles calculator HTTP requests
type Handler struct {
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
	router.HandleFunc("/add/{a}/{b}", h.handleAdd).Methods("GET")
	router.HandleFunc("/subtract/{a}/{b}", h.handleSubtract).Methods("GET")
	router.HandleFunc("/multiply/{a}/{b}", h.handleMultiply).Methods("GET")
	router.HandleFunc("/divide/{a}/{b}", h.handleDivide).Methods("GET")
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
// @Failure 400 {object} Response
// @Router /add/{a}/{b} [get]
func (h *Handler) handleAdd(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.writeError(w, err.Error())
		return
	}
	result := h.calc.Add(a, b)
	h.writeResult(w, result)
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
// @Failure 400 {object} Response
// @Router /subtract/{a}/{b} [get]
func (h *Handler) handleSubtract(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.writeError(w, err.Error())
		return
	}
	result := h.calc.Subtract(a, b)
	h.writeResult(w, result)
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
// @Failure 400 {object} Response
// @Router /multiply/{a}/{b} [get]
func (h *Handler) handleMultiply(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.writeError(w, err.Error())
		return
	}
	result := h.calc.Multiply(a, b)
	h.writeResult(w, result)
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
// @Failure 400 {object} Response
// @Router /divide/{a}/{b} [get]
func (h *Handler) handleDivide(w http.ResponseWriter, r *http.Request) {
	a, b, err := h.getNumbers(r)
	if err != nil {
		h.writeError(w, err.Error())
		return
	}
	result, err := h.calc.Divide(a, b)
	if err != nil {
		h.writeError(w, err.Error())
		return
	}
	h.writeResult(w, result)
}

// getNumbers extracts numbers from URL parameters
func (h *Handler) getNumbers(r *http.Request) (float64, float64, error) {
	vars := mux.Vars(r)

	a, err := strconv.ParseFloat(vars["a"], 64)
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.ParseFloat(vars["b"], 64)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

// writeResult writes a successful result
func (h *Handler) writeResult(w http.ResponseWriter, result float64) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: result})
}

// writeError writes an error response
func (h *Handler) writeError(w http.ResponseWriter, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response{Error: errMsg})
}
