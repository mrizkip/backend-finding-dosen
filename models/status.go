package models

import (
	"time"
)

type Status struct {
	ID              int       `db:"id" json:"id"`
	UserID          int       `db:"user_id" json:"user_id"`
	Status          string    `db:"status" json:"status"`
	DeskripsiStatus string    `db:"desc_status" json:"desc_status"`
	LastUpdate      time.Time `db:"last_update" json:"last_update"`
}

func NewStatus(userID int, status, deskripsiStatus, lastUpdate string) (*Status, error) {
	parsedDate, err := time.Parse("2006-01-02 15:04", lastUpdate)
	if err != nil {
		return nil, err
	}

	stats := &Status{
		UserID:          userID,
		Status:          status,
		DeskripsiStatus: deskripsiStatus,
		LastUpdate:      parsedDate,
	}

	if err := Dbm.Insert(stats); err != nil {
		return nil, err
	}

	return stats, err
}
