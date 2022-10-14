package auth

import (
	"encoding/json"
	"net/http"
	"project-golang/config"
	"project-golang/helpers"
	"project-golang/models"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil input Json dari Postman body
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
		// tambahkan return agar ketika ada error, codenya berhenti disini
	}

	defer r.Body.Close()

	// Check duplicate username
	var user models.User
	models.DB.Where("username= ?", userInput.Username).First(&user)

	lower := strings.ToLower(userInput.Username)

	if user.Username == lower {
		response := map[string]string{"message": "Duplicate Username"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Hash Password

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert database

	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func ReadOne(w http.ResponseWriter, r *http.Request) {
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

	var user models.User
	models.DB.First(&user, claims.Id)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var toko models.Toko
	models.DB.Where("user_id= ?", claims.UserId).First(&toko)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var barang []models.Barang
	models.DB.Where("id_toko = ?", toko.Id).Find(&barang)
	// models.DB.Raw(`SELECT * FROM barangs WHERE id_toko = ?`, toko.Id).Scan(&barang)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
	}

	user.Toko = toko
	user.Toko.Barang = barang

	helpers.ResponseJSON(w, http.StatusOK, user)
}

func ReadAnother(w http.ResponseWriter, r *http.Request) {
	// get id from params
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		response := map[string]string{"message": "id not found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var toko models.Toko
	// if err := models.DB.Where("user_id = ?", user.Id).First(&toko).Error; err != nil {
	// 	response := map[string]string{"message": "id not found"}
	// 	helpers.ResponseJSON(w, http.StatusBadRequest, response)
	// 	return
	// }
	models.DB.Where("user_id = ?", user.Id).First(&toko)

	var barang []models.Barang
	// if err := models.DB.Where("id_toko = ?", toko.Id).Find(&barang).Error; err != nil {
	// 	response := map[string]string{"message": "id not found"}
	// 	helpers.ResponseJSON(w, http.StatusBadRequest, response)
	// 	return
	// }
	models.DB.Where("id_toko = ?", toko.Id).Find(&barang)

	user.Password = ""
	user.Saldo = 0
	user.Toko = toko
	user.Toko.Barang = barang

	helpers.ResponseJSON(w, http.StatusOK, user)
}

func Update(w http.ResponseWriter, r *http.Request) {
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

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	// Check duplicate username
	var user models.User
	models.DB.Where("username= ?", userInput.Username).First(&user)

	if user.Username == userInput.Username {
		response := map[string]string{"message": "Duplicate Username"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Hash Password

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if models.DB.Where("id = ?", claims.Id).Updates(&userInput).RowsAffected == 0 {
		response := map[string]string{"message": "Id tidak ditemukan"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil input Json dari Postman body
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
		// tambahkan return agar ketika ada error, codenya berhenti disini
	}

	defer r.Body.Close()

	// check data di database

	var user models.User
	if err := models.DB.Where("username= ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return

		default:
			response := map[string]string{"message": err.Error()}
			helpers.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek password

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "password salah"}
		helpers.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan JWT
	// untuk menentukan waktu expired dari token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTclaim{
		Id:       int(user.Id),
		UserId:   int(user.Id),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan alogritma yang digunakan untuk login

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// set token ke cookies

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "berhasil login"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "berhasil logout"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func TutupAkun(w http.ResponseWriter, r *http.Request) {
	// get Id from token

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

	models.DB.Where("user_id= ?", claims.UserId).Delete(&models.Toko{})
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	models.DB.Delete(&models.User{}, claims.Id)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// delete cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "success delete data"}

	helpers.ResponseJSON(w, http.StatusOK, response)
}
