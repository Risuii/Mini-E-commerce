package ping

import (
	"net/http"
	"project-golang/helpers"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	helpers.ResponseJSON(w, http.StatusOK, "PING !!!")
}
