package toko

import (
	"encoding/json"
	"net/http"
	"project-golang/config"
	"project-golang/helpers"
	"project-golang/models"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func CreateToko(w http.ResponseWriter, r *http.Request) {

	// get userId from jwt

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

	// mengambil input Json dari Postman body
	var userInput models.Toko
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
		// tambahkan return agar ketika ada error, codenya berhenti disini
	}

	defer r.Body.Close()

	var namaToko models.Toko
	models.DB.Where("nama= ?", userInput.Nama).First(&namaToko)

	if namaToko.Nama == userInput.Nama {
		response := map[string]string{"message": "Duplicate Toko Name"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var toko models.Toko

	toko.Nama = userInput.Nama
	toko.Description = userInput.Description
	toko.UserId = int64(claims.Id)

	if err := models.DB.Create(&toko).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func ReadToko(w http.ResponseWriter, r *http.Request) {
	// Get User Id

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

	var toko []models.Toko
	// models.DB.Raw(`SELECT * FROM tokos WHERE tokos.user_id= ?`, claims.UserId).Scan(&toko)
	models.DB.Where("user_id= ?", claims.UserId).First(&toko)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, toko)
}

func ReadOneToko(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var toko models.Toko
	if err := models.DB.Where("id = ?", id).First(&toko).Error; err != nil {
		response := map[string]string{"message": "id not found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var barang []models.Barang
	if err := models.DB.Where("id_toko= ?", toko.Id).Find(&barang).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	toko.Barang = barang

	helpers.ResponseJSON(w, http.StatusOK, toko)
}

func EditToko(w http.ResponseWriter, r *http.Request) {
	// Get Id From Token

	c, err := r.Cookie("token")
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
	}

	tokenString := c.Value
	claims := &config.JWTclaim{}

	jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	// Get Input User
	var tokoInput models.Toko
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tokoInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	// check duplicate name
	var namaToko models.Toko
	models.DB.Where("nama= ?", tokoInput.Nama).First(&namaToko)

	if namaToko.Nama == tokoInput.Nama {
		response := map[string]string{"message": "Duplicate Toko Name"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if models.DB.Where("user_id = ?", claims.UserId).Updates(&tokoInput).RowsAffected == 0 {
		response := map[string]string{"message": "Id tidak ditemukan"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
