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
	fmt.Println("Start serving...")
	//register("POST", "/test", controller.Test, nil)
	register("POST", "/Trade/sell", controller.Sell, nil)
	register("POST", "/Trade/buy", controller.Buy, nil)
	register("POST", "/Trade/search", controller.Search, nil)
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
