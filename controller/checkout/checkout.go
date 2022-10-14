package checkout

import (
	"net/http"
	"project-golang/config"
	"project-golang/helpers"
	"project-golang/models"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	// get saldo from cookie

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

	var userSaldo models.User

	models.DB.Where("id = ?", claims.UserId).First(&userSaldo)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	var barang models.Barang

	models.DB.Where("id = ?", id).First(&barang)
	if err != nil {
		response := map[string]string{"message": "id not found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	if userSaldo.Saldo < barang.Harga || userSaldo.Saldo == 0 {
		response := map[string]string{"message": "Not enough saldo, please top up"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	newSaldo := userSaldo.Saldo - barang.Harga

	var user models.User

	user.Saldo = newSaldo

	models.DB.Where("id = ?", claims.UserId).Updates(&user)
	if err != nil {
		response := map[string]string{"message": "id not found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "success buy an item"}
	helpers.ResponseJSON(w, http.StatusOK, response)

}
