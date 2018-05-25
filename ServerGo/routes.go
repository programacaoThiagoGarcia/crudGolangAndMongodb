package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route model para criar rotas
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes devolve todas as rotas criadas
type Routes []Route

//NewRouter cria as rotas para enviar para a main
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/user",
		CreateUser,
	},
	Route{
		"GetUser",
		"GET",
		"/user",
		GetUsers,
	},
	Route{
		"GetUserName",
		"GET",
		"/userName/{name}",
		GetUserName,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user",
		DeleteUser,
	},
}
