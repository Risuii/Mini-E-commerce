package main

import (
	"project-golang/models"
	"project-golang/routes"
)

// var DB *gorm.DB

func main() {
	models.ConnectionDatabase()
	routes.Routing()
}
