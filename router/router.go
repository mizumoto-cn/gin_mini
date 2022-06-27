package router

import (
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

// HandlerFunc defines the request handler function used by ginmini.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// NewRouter returns a new router.
func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) AddRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	r.handlers[key] = handler
}
