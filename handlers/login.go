package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrizkip/backend-finding-dosen/env"
	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Message      string      `json:"message"`
	Token        string      `json:"token"`
	LoggedInUser models.User `json:"logged_in_user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req loginRequest

	if err := decoder.Decode(&req); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	var user models.User
	query := "select * from users where email=?"

	if err := models.Dbm.SelectOne(&user, query, req.Email); err != nil {
		errors.NewError("user not found", http.StatusUnauthorized).WriteTo(w)
		return
	}

	if err := user.VerifyPassword(req.Password); err != nil {
		errors.NewError("password incorrect", http.StatusUnauthorized).WriteTo(w)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 720).Unix(),
	})

	secret := env.Getenv("SECRET_KEY", "secret")
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		errors.NewError("can't sign your token", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(loginResponse{
		Message:      "logged in",
		Token:        tokenString,
		LoggedInUser: user,
	})
}
