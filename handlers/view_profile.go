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
	UserID int    `json:"user_id"`
	Nama   string `json:"nama"`
	Status string `json:"status"`
}

// DosenResponse represent json response individual dosen
type DosenResponse struct {
	UserID          int    `json:"user_id"`
	Nama            string `json:"nama"`
	Email           string `json:"email"`
	JenisIdentitas  string `json:"jenis_identitas"`
	NoIdentitas     string `json:"no_identitas"`
	NoTelpon        string `json:"no_Telpon"`
	Status          string `json:"status"`
	DeskripsiStatus string `json:"desc_status"`
	KetStatus       string `json:"ket_status"`
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

	query := `SELECT users.id as UserID, users.nama, status.status
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

// FetchDosenByID represent a request for get dosen profile by id dosen
func FetchDosenByID(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "id")

	dosen, err := fetchDosen(userID)

	if err != nil {
		errors.NewError("Can't fetch profile", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(dosen)
}

func fetchDosen(id string) (DosenResponse, error) {
	var dosen DosenResponse

	query := `
	SELECT users.id as UserID, users.nama, users.email, users.jenis_identitas as JenisIdentitas, 
	users.no_identitas as NoIdentitas, users.no_telpon as NoTelpon, 
	status.status, status.desc_status as DeskripsiStatus, status.ket_status as KetStatus
	FROM users JOIN status
	WHERE users.id = ? AND users.id = status.user_id;
	`
	if err := models.Dbm.SelectOne(&dosen, query, id); err != nil {
		return DosenResponse{}, err
	}

	return dosen, nil
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
