package middleware

import (
	"net/http"
	"project-golang/config"
	"project-golang/helpers"

	"github.com/golang-jwt/jwt/v4"
)

// untuk check apakah yang akses sudah punya token jwt atau belom (sudah login atau belom)
func JWTmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		// mengambil token

		tokenString := c.Value
		claims := &config.JWTclaim{}

		// Parsing token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		// check token masih bisa atau sudah expired

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return

			case jwt.ValidationErrorExpired:
				// token expired
				response := map[string]string{"message": "Unauthorized because token expired"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return

			default:
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}

		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
