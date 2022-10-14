package toko

import (
	"project-golang/controller/toko"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func RoutingToko() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	api.HandleFunc("/create", toko.CreateToko).Methods("POST")
	api.HandleFunc("/toko", toko.ReadToko).Methods("GET")
	api.HandleFunc("/update-toko", toko.EditToko).Methods("PATCH")
}
