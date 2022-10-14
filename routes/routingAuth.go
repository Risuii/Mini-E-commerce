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
	r.HandleFunc("/profile/{id}", auth.ReadAnother).Methods("GET")

	// User

	api.HandleFunc("/update", auth.Update).Methods("PATCH")
	api.HandleFunc("/profile", auth.ReadOne).Methods("GET")
	api.HandleFunc("/delete-akun", auth.TutupAkun).Methods("DELETE")

	// Toko

	api.HandleFunc("/create", toko.CreateToko).Methods("POST")
	api.HandleFunc("/toko", toko.ReadToko).Methods("GET")
	r.HandleFunc("/toko/{id}", toko.ReadOneToko).Methods("GET")
	api.HandleFunc("/update-toko", toko.EditToko).Methods("PATCH")

	// Product Owner

	api.HandleFunc("/input", product.InputBarang).Methods("POST")
	api.HandleFunc("/barang", product.ReadBarang).Methods("GET")
	api.HandleFunc("/barang/{id}", product.EditBarang).Methods("PATCH")
	r.HandleFunc("/barang/{id}", product.Readone).Methods("GET")
	api.HandleFunc("/delete-barang/{id}", product.DeleteBarang).Methods("DELETE")

	// Payment
	api.HandleFunc("/topup", checkout.Topup).Methods("POST")
	api.HandleFunc("/checkout/{id}", checkout.Checkout).Methods("POST")

	fmt.Println("PORT: 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
