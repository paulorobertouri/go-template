package common

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler defines the interface for HTTP handlers
type Handler interface {
	RegisterRoutes(router *mux.Router)
}

// Route represents a single route configuration
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
	Name    string // Optional name for the route
}

// RouteGroup represents a group of routes with a common prefix
type RouteGroup struct {
	Prefix string
	Routes []Route
}

// RegisterGroup registers a group of routes with a common prefix
func RegisterGroup(router *mux.Router, group RouteGroup) {
	for _, route := range group.Routes {
		path := group.Prefix + route.Path
		r := router.HandleFunc(path, route.Handler).Methods(route.Method)
		if route.Name != "" {
			r.Name(route.Name)
		}
	}
}

// SimpleRoute creates a simple route with just path, method and handler
func SimpleRoute(path, method string, handler http.HandlerFunc) Route {
	return Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}

// NamedRoute creates a route with a name
func NamedRoute(path, method, name string, handler http.HandlerFunc) Route {
	return Route{
		Path:    path,
		Method:  method,
		Handler: handler,
		Name:    name,
	}
}
