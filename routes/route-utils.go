package routes

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Pattern string
	Handler gin.HandlerFunc
	// Middleware mux.MiddlewareFunc
}

var routes []Route

func register(method, pattern string, handler gin.HandlerFunc) {
	routes = append(routes, Route{method, pattern, handler})
}

// Bind HandlerFunc to Routes
func NewRouter() *gin.Engine {
	r := gin.Default()
	for _, route := range routes {
		r.Handle(route.Method, route.Pattern, route.Handler)
	}

	return r
}
