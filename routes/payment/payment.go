package payment

import (
	"project-golang/controller/checkout"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func RoutingPayment() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	api.HandleFunc("/topup", checkout.Topup).Methods("POST")
	api.HandleFunc("/checkout/{id}", checkout.Checkout).Methods("POST")
}
