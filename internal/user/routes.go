package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	usersPath        = "/users"
	usersIDPath      = "/users/{id}"
	invalidUserIDMsg = "invalid user ID"
)

// CreateUserRequest represents a user creation request
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler handles user HTTP requests
type Handler struct {
	service *Service
}

// NewHandler creates a new user handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterRoutes registers user routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc(usersPath, h.handleGetAllUsers).Methods("GET")
	router.HandleFunc(usersIDPath, h.handleGetUser).Methods("GET")
	router.HandleFunc(usersPath, h.handleCreateUser).Methods("POST")
	router.HandleFunc(usersIDPath, h.handleUpdateUser).Methods("PUT")
	router.HandleFunc(usersIDPath, h.handleDeleteUser).Methods("DELETE")
}

// handleGetAllUsers handles GET requests for all users
// @Summary Get all users
// @Description Retrieves a list of all users in the system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAllUsers()
	h.writeJSON(w, users)
}

// handleGetUser handles GET requests for a specific user
// @Summary Get user by ID
// @Description Retrieves a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDFromURL(r)
	if err != nil {
		h.writeError(w, invalidUserIDMsg, http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			h.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			h.writeError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	h.writeJSON(w, user)
}

// handleCreateUser handles POST requests to create a new user
// @Summary Create a new user
// @Description Creates a new user with the provided name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User creation request"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func (h *Handler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	h.writeJSON(w, user)
}

// handleUpdateUser handles PUT requests to update a user
// @Summary Update user by ID
// @Description Updates an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body CreateUserRequest true "User update request"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDFromURL(r)
	if err != nil {
		h.writeError(w, invalidUserIDMsg, http.StatusBadRequest)
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.service.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			h.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			h.writeError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	h.writeJSON(w, user)
}

// handleDeleteUser handles DELETE requests to delete a user
// @Summary Delete user by ID
// @Description Deletes a user from the system
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDFromURL(r)
	if err != nil {
		h.writeError(w, invalidUserIDMsg, http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			h.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			h.writeError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getIDFromURL extracts the ID from URL parameters
func (h *Handler) getIDFromURL(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}

// writeJSON writes a JSON response
func (h *Handler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// writeError writes an error response
func (h *Handler) writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
