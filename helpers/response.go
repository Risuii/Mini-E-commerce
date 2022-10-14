package helpers

import (
	"encoding/json"
	"net/http"
)

// fungsi untuk response
// code disini untuk http status
// interface agar bisa menerima semua bentuk data

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
