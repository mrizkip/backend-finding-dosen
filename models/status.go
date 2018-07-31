package models

import (
	"time"
)

// Status object
type Status struct {
	ID               int       `db:"id" json:"id"`
	UserID           int       `db:"user_id" json:"user_id"`
	Status           string    `db:"status" json:"status"`
	DeskripsiStatus  string    `db:"desc_status" json:"desc_status"`
	Posisi           string    `db:"posisi" json:"posisi"`
	KeteranganStatus string    `db:"ket_status" json:"ket_status"`
	LastUpdate       time.Time `db:"last_update" json:"last_update"`
}

// NewStatus represent new statur for new  dosen registered
func NewStatus(userID int, status, deskripsiStatus, keteranganStatus, lastUpdate string) (*Status, error) {
	parsedDate, err := time.Parse("2006-01-02 15:04", lastUpdate)
	if err != nil {
		return nil, err
	}

	stats := &Status{
		UserID:           userID,
		Status:           status,
		DeskripsiStatus:  deskripsiStatus,
		KeteranganStatus: keteranganStatus,
		LastUpdate:       parsedDate,
	}

	if err := Dbm.Insert(stats); err != nil {
		return nil, err
	}

	return stats, err
}
