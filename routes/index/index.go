package index

import (
	"project-golang/controller/home"
	"project-golang/controller/ping"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func RoutingIndex() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	r.HandleFunc("/", home.Home).Methods("GET")
	r.HandleFunc("/ping", ping.Ping).Methods("GET")
}
