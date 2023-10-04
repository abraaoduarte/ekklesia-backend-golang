package server

import (
	"github.com/abraaoduarte/ekklesia-backend-golang/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Optionally, you can set up some global middleware here

	return router
}

func RegisterHandlers(router *mux.Router) {
	subRouter := router.PathPrefix("/v1").Subrouter()
	subRouter.HandleFunc("/user/{id}", handlers.GetUserDetail).Methods("GET")
}
