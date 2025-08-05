package common

import (
	"encoding/json"
	"net/http"
)

// BaseHandler provides common functionality for all handlers
type BaseHandler struct{}

// ParseJSON parses JSON from request body into the given interface
func (h *BaseHandler) ParseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// GetURLParams extracts URL parameters from the request
func (h *BaseHandler) GetURLParams(r *http.Request) *URLParams {
	return GetURLParams(r)
}

// WriteSuccess writes a successful response
func (h *BaseHandler) WriteSuccess(w http.ResponseWriter, data interface{}) {
	WriteSuccess(w, data)
}

// WriteCreated writes a created response
func (h *BaseHandler) WriteCreated(w http.ResponseWriter, data interface{}) {
	WriteCreated(w, data)
}

// WriteError writes an error response
func (h *BaseHandler) WriteError(w http.ResponseWriter, statusCode int, message string) {
	WriteError(w, statusCode, message)
}

// WriteBadRequest writes a bad request error
func (h *BaseHandler) WriteBadRequest(w http.ResponseWriter, message string) {
	WriteBadRequest(w, message)
}

// WriteNotFound writes a not found error
func (h *BaseHandler) WriteNotFound(w http.ResponseWriter, message string) {
	WriteNotFound(w, message)
}

// WriteInternalError writes an internal server error
func (h *BaseHandler) WriteInternalError(w http.ResponseWriter, message string) {
	WriteInternalError(w, message)
}
