package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type registerRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	Nama           string `json:"nama"`
	JenisIdentitas string `json:"jenis_identitas"`
	NoIdentitas    string `json:"no_identitas"`
}

type registerResponse struct {
	Message string `json:"message"`
}

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

	if _, err := models.NewUser(reqData.Email, reqData.Password, reqData.Nama, reqData.JenisIdentitas, reqData.NoIdentitas, role); err != nil {
		errors.NewError("user already registered", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(registerResponse{
		Message: "user registered",
	})
}
