package routes

import (
	"fmt"
	"log"
	"net/http"
	"project-golang/controller/auth"
	"project-golang/controller/checkout"
	"project-golang/controller/home"
	"project-golang/controller/ping"
	"project-golang/controller/product"
	"project-golang/controller/toko"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func Routing() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	// Home
	r.HandleFunc("/", home.Home).Methods("GET")

	// Ping

	r.HandleFunc("/ping", ping.Ping).Methods("GET")

	// Auth

	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET")
	r.HandleFunc("/profile/{id}", auth.ReadAnotherProfile).Methods("GET")
	r.HandleFunc("/store/{id}", toko.ReadOneToko).Methods("GET")
	r.HandleFunc("/item/{id}", product.Readone).Methods("GET")

	// User

	api.HandleFunc("/update", auth.Update).Methods("PATCH")
	api.HandleFunc("/profile", auth.ReadOne).Methods("GET")
	api.HandleFunc("/delete-profile", auth.DeleteProfile).Methods("DELETE")

	// Toko

	api.HandleFunc("/create", toko.CreateToko).Methods("POST")
	api.HandleFunc("/store", toko.ReadToko).Methods("GET")
	api.HandleFunc("/update-toko", toko.EditToko).Methods("PATCH")

	// Product Owner

	api.HandleFunc("/input", product.InputBarang).Methods("POST")
	api.HandleFunc("/item", product.ReadBarang).Methods("GET")
	api.HandleFunc("/item/{id}", product.EditBarang).Methods("PATCH")
	api.HandleFunc("/delete-item/{id}", product.DeleteBarang).Methods("DELETE")

	// Payment

	api.HandleFunc("/topup", checkout.Topup).Methods("POST")
	api.HandleFunc("/checkout/{id}", checkout.Checkout).Methods("POST")

	fmt.Println("PORT: 8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}
