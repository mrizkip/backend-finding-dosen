package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type changePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type changeProfileRequest struct {
	Email string `json:"email"`
	Nama  string `json:"full_name"`
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req changeProfileRequest
	if err := decoder.Decode(&req); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	userId := r.Context().Value("user_id").(int)

	var user models.User
	if err := models.Dbm.SelectOne(&user, "select * from users where id=?", userId); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}

	user.Nama = req.Nama

	if _, err := models.Dbm.Update(&user); err != nil {
		errors.NewError("can't change profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "profile updated",
	})
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req changePasswordRequest
	if err := decoder.Decode(&req); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	userId := r.Context().Value("user_id").(int)

	var user models.User
	if err := models.Dbm.SelectOne(&user, "select * from users where id=?", userId); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}

	if err := user.VerifyPassword(req.OldPassword); err != nil {
		errors.NewError("incorrect password", http.StatusInternalServerError).WriteTo(w)
		return
	}

	user.HashPassword(req.NewPassword)

	if _, err := models.Dbm.Update(&user); err != nil {
		errors.NewError("can't change password", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "password updated",
	})
}
