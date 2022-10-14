package main

import (
	"project-golang/models"
	"project-golang/routes"
)

func main() {
	models.ConnectionDatabase()
	routes.Routing()
}
