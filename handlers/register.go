package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type registerRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	Nama           string `json:"nama"`
	JenisIdentitas string `json:"jenis_identitas"`
	NoIdentitas    string `json:"no_identitas"`
	NoTelpon       string `json:"no_telpon"`
}

type registerResponse struct {
	Message string `json:"message"`
}

// Register represent a request for register user
func Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqData registerRequest
	if err := decoder.Decode(&reqData); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	var role string

	if (reqData.JenisIdentitas == "NIP") || (reqData.JenisIdentitas == "NIK") {
		role = "dosen"
	} else {
		role = "mahasiswa"
	}

	if _, err := models.NewUser(reqData.Email, reqData.Password, reqData.Nama, reqData.JenisIdentitas, reqData.NoIdentitas, reqData.NoTelpon, role); err != nil {
		errors.NewError("user already registered", http.StatusInternalServerError).WriteTo(w)
		return
	}

	if role == "dosen" {
		lastUpdate := time.Now()
		formatedLastUpdate := lastUpdate.Format("2006-01-02 15:04")

		var user models.User

		query := "SELECT * FROM users WHERE email = ?"
		if err := models.Dbm.SelectOne(&user, query, reqData.Email); err != nil {
			errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
			return
		}

		if _, err := models.NewStatus(user.ID, "Tidak Aktif", "", formatedLastUpdate); err != nil {
			errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
			return
		}
	}

	json.NewEncoder(w).Encode(registerResponse{
		Message: "user registered",
	})
}
