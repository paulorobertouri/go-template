package common

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ErrParameterNotFound = errors.New("parameter not found")
)

// URLParams provides helper methods to extract parameters from URL
type URLParams struct {
	vars map[string]string
}

// NewURLParams creates a new URLParams from mux.Vars
func NewURLParams(vars map[string]string) *URLParams {
	return &URLParams{vars: vars}
}

// String returns a string parameter
func (p *URLParams) String(key string) (string, bool) {
	value, exists := p.vars[key]
	return value, exists
}

// Int returns an integer parameter
func (p *URLParams) Int(key string) (int, error) {
	value, exists := p.vars[key]
	if !exists {
		return 0, ErrParameterNotFound
	}
	return strconv.Atoi(value)
}

// Float64 returns a float64 parameter
func (p *URLParams) Float64(key string) (float64, error) {
	value, exists := p.vars[key]
	if !exists {
		return 0, ErrParameterNotFound
	}
	return strconv.ParseFloat(value, 64)
}

// GetURLParams extracts URL parameters from a request using mux.Vars
func GetURLParams(r *http.Request) *URLParams {
	return NewURLParams(mux.Vars(r))
}
