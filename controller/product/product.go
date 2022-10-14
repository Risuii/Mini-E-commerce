package product

import (
	"net/http"
	"project-golang/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "kemeja",
			"stok":         100,
		},
		{
			"id":           2,
			"nama_product": "celana jeans",
			"stok":         100,
		},
		{
			"id":           3,
			"nama_product": "kacamata",
			"stok":         100,
		},
	}

	helpers.ResponseJSON(w, http.StatusOK, data)
}
