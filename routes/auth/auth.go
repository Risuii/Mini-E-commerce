package auth

import (
	"project-golang/controller/auth"
	"project-golang/controller/product"
	"project-golang/controller/toko"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func RoutingAuth() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET")
	r.HandleFunc("/profile/{id}", auth.ReadAnotherProfile).Methods("GET")
	r.HandleFunc("/toko/{id}", toko.ReadOneToko).Methods("GET")
	r.HandleFunc("/barang/{id}", product.Readone).Methods("GET")
}
