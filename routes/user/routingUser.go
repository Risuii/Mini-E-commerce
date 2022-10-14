package user

import (
	"project-golang/controller/auth"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func RoutingUser() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	api.HandleFunc("/update", auth.Update).Methods("PATCH")
	api.HandleFunc("/profile", auth.ReadOne).Methods("GET")
	api.HandleFunc("/delete-akun", auth.DeleteProfile).Methods("DELETE")
}
