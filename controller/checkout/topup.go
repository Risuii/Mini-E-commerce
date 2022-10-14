package checkout

import (
	"encoding/json"
	"net/http"
	"project-golang/config"
	"project-golang/helpers"
	"project-golang/models"

	"github.com/golang-jwt/jwt"
)

func Topup(w http.ResponseWriter, r *http.Request) {
	// get id from cookie token

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
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var oldSaldo models.User
	models.DB.Where("id = ?", claims.UserId).First(&oldSaldo)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var user models.User

	NewSaldo := oldSaldo.Saldo + userInput.Saldo

	user.Saldo = NewSaldo

	models.DB.Where("id = ?", claims.UserId).Updates(&user)
	if err != nil {
		response := map[string]string{"message": "id not found"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "success top up"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
