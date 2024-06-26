package routes

import (
	"net/http"
	"petopia-server/controller"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("POST", "/api/todo", controller.AddTodo, nil)
	register("GET", "/api/todo/{id}", controller.GetTodoById, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		r.Methods(route.Method).Path(route.Pattern).Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method string, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
