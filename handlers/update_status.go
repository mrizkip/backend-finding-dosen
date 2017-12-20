package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

func UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID          int    `json:"user_id"`
		Status          string `json:"status"`
		DeskripsiStatus string `json:"desc_status"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	lastUpdate := time.Now()
	formatedLastUpdate := lastUpdate.Format("2006-01-02 15:04")

	userID := r.Context().Value("user_id").(int)
	var stats models.Status
	query := "SELECT * FROM status WHERE user_id = ?"
	if err := models.Dbm.SelectOne(&stats, query, userID); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}

	parsedDate, err := time.Parse("2006-01-02 15:04", formatedLastUpdate)
	if err != nil {
		return
	}

	stats.Status = req.Status
	stats.DeskripsiStatus = req.DeskripsiStatus
	stats.LastUpdate = parsedDate

	if _, err := models.Dbm.Update(&stats); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "update status successfully",
	})

}
