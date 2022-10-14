package main

import (
	"project-golang/models"
	"project-golang/routes"
)

func main() {
	models.ConnectionDatabase()

	routes.Routing()

	// index.RoutingIndex()
	// auth.RoutingAuth()
	// user.RoutingUser()
	// toko.RoutingToko()
	// productOwner.RoutingProductOwner()
	// payment.RoutingPayment()

	// fmt.Println("PORT: 8080")

	// r := mux.NewRouter()

	// log.Fatal(http.ListenAndServe(":8080", r))
}
