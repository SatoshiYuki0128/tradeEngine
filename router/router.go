package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"tradeEngine/controller"
	"tradeEngine/model"
)

var routes []model.Route

func init() {
	fmt.Println("Route Init")
	register("POST", "/test", controller.Test, nil)
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	handler := cors.Default().Handler(r)
	return handler
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, model.Route{Method: method, Pattern: pattern, Handler: handler, Middleware: middleware})
}
