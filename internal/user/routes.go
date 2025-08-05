package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulorobertouri/go-template/internal/common"
)

// CreateUserRequest represents a user creation request
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Handler handles user HTTP requests
type Handler struct {
	common.BaseHandler
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
	routes := common.RouteGroup{
		Prefix: "/users",
		Routes: []common.Route{
			common.NamedRoute("", "GET", "users.getAll", h.handleGetAllUsers),
			common.NamedRoute("", "POST", "users.create", h.handleCreateUser),
			common.NamedRoute("/{id}", "GET", "users.get", h.handleGetUser),
			common.NamedRoute("/{id}", "PUT", "users.update", h.handleUpdateUser),
			common.NamedRoute("/{id}", "DELETE", "users.delete", h.handleDeleteUser),
		},
	}
	common.RegisterGroup(router, routes)
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
	h.WriteSuccess(w, users)
}

// handleGetUser handles GET requests for a specific user
// @Summary Get user by ID
// @Description Retrieves a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Router /users/{id} [get]
func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDFromURL(r)
	if err != nil {
		h.WriteBadRequest(w, "invalid user ID")
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			h.WriteNotFound(w, err.Error())
		} else {
			h.WriteBadRequest(w, err.Error())
		}
		return
	}

	h.WriteSuccess(w, user)
}

// handleCreateUser handles POST requests to create a new user
// @Summary Create a new user
// @Description Creates a new user with the provided name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User creation request"
// @Success 201 {object} User
// @Failure 400 {object} common.ErrorResponse
// @Router /users [post]
func (h *Handler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := h.ParseJSON(r, &req); err != nil {
		h.WriteBadRequest(w, "invalid JSON payload")
		return
	}

	user, err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		h.WriteBadRequest(w, err.Error())
		return
	}

	h.WriteCreated(w, user)
}

// handleUpdateUser handles PUT requests to update an existing user
// @Summary Update user by ID
// @Description Updates an existing user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body CreateUserRequest true "User update request"
// @Success 200 {object} User
// @Failure 400 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Router /users/{id} [put]
func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDFromURL(r)
	if err != nil {
		h.WriteBadRequest(w, "invalid user ID")
		return
	}

	var req CreateUserRequest
	if err := h.ParseJSON(r, &req); err != nil {
		h.WriteBadRequest(w, "invalid JSON payload")
		return
	}

	user, err := h.service.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			h.WriteNotFound(w, err.Error())
		} else {
			h.WriteBadRequest(w, err.Error())
		}
		return
	}

	h.WriteSuccess(w, user)
}

// handleDeleteUser handles DELETE requests to delete a user
// @Summary Delete user by ID
// @Description Deletes a user from the system
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDFromURL(r)
	if err != nil {
		h.WriteBadRequest(w, "invalid user ID")
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			h.WriteNotFound(w, err.Error())
		} else {
			h.WriteBadRequest(w, err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getIDFromURL extracts the ID parameter from the URL
func (h *Handler) getIDFromURL(r *http.Request) (int, error) {
	params := h.GetURLParams(r)
	return params.Int("id")
}
