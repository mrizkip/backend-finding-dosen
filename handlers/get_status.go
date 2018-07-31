package handlers

import (
	"net/http"
	"time"

	"encoding/json"

	"goji.io/pat"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

// StatusLocation represent json response for get position of dosen
type StatusLocation struct {
	UserID     int       `json:"user_id"`
	Nama       string    `json:"nama"`
	Posisi     string    `json:"posisi"`
	LastUpdate time.Time `json:"last_update"`
}

// GetDosenLocation represent method for get location of dosen
func GetDosenLocation(w http.ResponseWriter, r *http.Request) {
	var lokasi StatusLocation
	userID := pat.Param(r, "id")

	query := `SELECT users.id as UserID, users.nama, status.posisi, status.last_update as LastUpdate
	FROM users JOIN status
	WHERE users.id = ? AND users.id = status.user_id;`

	if err := models.Dbm.SelectOne(&lokasi, query, userID); err != nil {
		errors.NewError("Can't fetch location", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(lokasi)
}
