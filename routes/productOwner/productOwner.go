package productOwner

import (
	"project-golang/controller/product"
	"project-golang/middleware"

	"github.com/gorilla/mux"
)

func RoutingProductOwner() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTmiddleware)

	api.HandleFunc("/input", product.InputBarang).Methods("POST")
	api.HandleFunc("/barang", product.ReadBarang).Methods("GET")
	api.HandleFunc("/barang/{id}", product.EditBarang).Methods("PATCH")
	api.HandleFunc("/delete-barang/{id}", product.DeleteBarang).Methods("DELETE")
}
