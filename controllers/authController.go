package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/rokibulislam/gocommerce/domain"
)

// func GenerateToken(user domain.User) (string, error) {
// 	var err domain.Error
// 	// secret := "secret"
// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
// 		"email": user.Email,
// 	})

// 	return token, err
// }

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user domain.User
		var error domain.Error
		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing"
			// responseWithError()
		}

		if user.Password == "" {
			error.Message = "password is missing"
			// responseWithError()
		}
	}
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
