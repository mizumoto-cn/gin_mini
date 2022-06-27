package ginmini

import (
	"log"
	"net/http"

	"github.com/mizumoto-cn/ginmini/router"
)

// Engine is the framework's instance.
type Engine struct {
	// router is the map of request handlers.
	// key is the method + "-" + path.
	router *router.Router
}

// New creates a new ginmini engine.
func New() *Engine {
	return &Engine{router: router.NewRouter()}
}

// addRoute adds a new route to the engine.
func (e *Engine) addRoute(method string, path string, handler router.HandlerFunc) {
	log.Printf("Route %4s - %s", method, path)
	e.router.AddRoute(method, path, handler)
}

// GET defines a method to add a GET route.
func (e *Engine) GET(path string, handler router.HandlerFunc) {
	e.addRoute("GET", path, handler)
}

// POST defines a method to add a POST route.
func (e *Engine) POST(path string, handler router.HandlerFunc) {
	e.addRoute("POST", path, handler)
}

// RUN starts the server. It's a shorthand for http.ListenAndServe(addr, engine).
func (e *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP serves the incoming requests.
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
