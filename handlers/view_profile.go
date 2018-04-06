package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"goji.io/pat"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

// DosenResponseAll represent json response when fetch all dosen
type DosenResponseAll struct {
	ID     int    `json: id`
	Nama   string `json: nama`
	Status string `json: status`
}

// FetchMyProfile represent a request for get my profile
func FetchMyProfile(w http.ResponseWriter, r *http.Request) {
	myID := r.Context().Value("user_id").(int)

	user, err := fetchUser(strconv.Itoa(myID))

	if err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]models.User{
		"data": user,
	})
}

// FetchUserProfileByID represent a requset for get user profile by ID
func FetchUserProfileByID(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "id")

	user, err := fetchUser(userID)

	if err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]models.User{
		"data": user,
	})
}

// FetchAllDosenProfile represent a reuqest for get all dosen profile
func FetchAllDosenProfile(w http.ResponseWriter, r *http.Request) {
	var users []DosenResponseAll

	role := "dosen"

	query := `SELECT users.id, users.nama, status.status
	FROM users JOIN status
	WHERE role=? AND users.id = status.user_id;`

	if _, err := models.Dbm.Select(&users, query, role); err != nil {
		errors.NewError("can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string][]DosenResponseAll{
		"data": users,
	})
}

// FetchUserProfileByEmail represent a request for get user profile by email
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
