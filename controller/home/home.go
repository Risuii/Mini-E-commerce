package home

import (
	"net/http"
	"project-golang/helpers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helpers.ResponseJSON(w, http.StatusOK, "Welcome to mini e-commerce service")
}
