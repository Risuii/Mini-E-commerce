package product

import (
	"encoding/json"
	"net/http"
	"project-golang/config"
	"project-golang/helpers"
	"project-golang/models"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func InputBarang(w http.ResponseWriter, r *http.Request) {
	// get token from cookie

	c, err := r.Cookie("token")
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	tokenString := c.Value
	claims := &config.JWTclaim{}

	jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	var toko models.Toko
	models.DB.Where("user_id= ?", claims.UserId).First(&toko)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// get body postman

	var InputBarang models.Barang
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&InputBarang); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	var Barang models.Barang

	Barang.IdToko = int(toko.Id)
	Barang.Nama = InputBarang.Nama
	Barang.Description = InputBarang.Description
	Barang.Stock = InputBarang.Stock
	Barang.Harga = InputBarang.Harga

	models.DB.Create(&Barang)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func ReadBarang(w http.ResponseWriter, r *http.Request) {
	// get id from token
	c, err := r.Cookie("token")
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	tokenString := c.Value
	claims := &config.JWTclaim{}

	jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	var toko models.Toko
	models.DB.Where("user_id= ?", claims.UserId).First(&toko)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var barang []models.Barang
	models.DB.Where("id_toko= ?", toko.Id).Find(&barang)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, barang)
}

func Readone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var barang models.Barang

	// pakai .Error agar kalau idnya ketemu tidak error
	if err := models.DB.Where("id = ?", id).First(&barang).Error; err != nil {
		response := map[string]string{"message": "Id not found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, barang)
}

func EditBarang(w http.ResponseWriter, r *http.Request) {
	// input from body

	var userInput models.Barang
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	params := mux.Vars(r)
	id := params["id"]

	if models.DB.Where("id = ?", id).Updates(&userInput).RowsAffected == 0 {
		response := map[string]string{"message": "Id Not Found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if models.DB.Where("id = ?", id).Delete(&models.Barang{}).RowsAffected == 0 {
		response := map[string]string{"message": "Id Not Found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
	}

	response := map[string]string{"message": "success deleted data"}
	helpers.ResponseJSON(w, http.StatusOK, response)

}
