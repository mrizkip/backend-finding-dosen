package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"goji.io/pat"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

func FetchMyProfile(w http.ResponseWriter, r *http.Request) {
	myId := r.Context().Value("user_id").(int)

	user, err := fetchUser(strconv.Itoa(myId))

	if err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]models.User{
		"data": user,
	})
}

func FetchUserProfile(w http.ResponseWriter, r *http.Request) {
	userId := pat.Param(r, "id")

	user, err := fetchUser(userId)

	fmt.Println(userId)

	if err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]models.User{
		"data": user,
	})
}

func FetchAllDosenProfile(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	role := "dosen"

	query := `SELECT id, email, nama, jenis_identitas, no_identitas, role
	FROM users
	WHERE role=?
	`

	if _, err := models.Dbm.Select(&users, query, role); err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.User{
		"data": users,
	})
}

func FetchUserProfileByEmail(w http.ResponseWriter, r *http.Request) {
	email := pat.Param(r, "email")

	var user models.User

	query := `
	SELECT id, email, nama, jenis_identitas, no_identitas, role
	FROM users
	WHERE email=?
	`
	if err := models.Dbm.SelectOne(&user, query, email); err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]models.User{
		"data": user,
	})
}

func fetchUser(id string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, email, nama, jenis_identitas, no_identitas, role
	FROM users
	WHERE id=?
	`
	if err := models.Dbm.SelectOne(&user, query, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}
