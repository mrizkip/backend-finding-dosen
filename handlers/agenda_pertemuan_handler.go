package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type agendaPertemuanRequest struct {
	UserID           int       `json:"user_id"`
	DosenID          int       `json:"dosen_id"`
	Judul            string    `json:"judul"`
	Keterangan       string    `json:"keterangan"`
	TanggalPertemuan time.Time `json:"tanggal_pertemuan"`
}

// AddAgendaPertemuan represent a request for add new agenda pertemuan
func AddAgendaPertemuan(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqData agendaPertemuanRequest
	if err := decoder.Decode(&reqData); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	tanggalPertemuan := reqData.TanggalPertemuan.Format("2006-01-02 15:04")

	if _, err := models.NewAgendaPertemuan(reqData.UserID, reqData.DosenID, reqData.Judul, reqData.Keterangan, tanggalPertemuan); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}
}

// FetchMyAgendaPertemuan represetnt a request for get Agenda pertemuan from user
func FetchMyAgendaPertemuan(w http.ResponseWriter, r *http.Request) {
	var agenda []models.AgendaPertemuan

	myID := r.Context().Value("user_id").(int)

	query := `SELECT * FROM agenda_pertemuan
	WHERE user_id = ?
	`

	if _, err := models.Dbm.Select(&agenda, query, myID); err != nil {
		errors.NewError("Tidak ada Agenda Pertemuan", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.AgendaPertemuan{
		"data": agenda,
	})
}

// FetchAgendaPertemuanByDosenID represent a request to get list of agenda pertemuan from dosen
func FetchAgendaPertemuanByDosenID(w http.ResponseWriter, r *http.Request) {
	var agenda []models.AgendaPertemuan

	myID := r.Context().Value("user_id").(int)

	query := `SELECT * FROM agenda_pertemuan
	WHERE dosen_id = ?
	`

	if _, err := models.Dbm.Select(&agenda, query, myID); err != nil {
		errors.NewError("Tidak ada Agenda Pertemuan", http.StatusInternalServerError).WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.AgendaPertemuan{
		"data": agenda,
	})
}
