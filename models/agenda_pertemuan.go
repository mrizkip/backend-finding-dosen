package models

import (
	"time"
)

type AgendaPertemuan struct {
	ID               int       `db:"id" json:"id"`
	UserID           int       `db:"user_id" json:"user_id"`
	DosenID          int       `db:"dosen_id" json:"dosen_id"`
	Status           string    `db:"status" json:"status"`
	Judul            string    `db:"judul" json:"judul"`
	Keterangan       string    `db:"keterangan" json:"keterangan"`
	TanggalPertemuan time.Time `db:"tanggal_pertemuan" json:"tanggal_pertemuan"`
	TanggalDibuat    time.Time `db:"tanggal_dibuat" json:"tanggal_dibuat"`
}

type CatatanAgendaPertemuan struct {
	ID                int       `db:"id" json:"id"`
	AgendaPertemuanID int       `db:"agenda_pertemuan_id" json:"agenda_pertemuan_id"`
	UserID            int       `db:"user_id" json:"user_id"`
	Catatan           string    `db:"catatan" json:"catatan"`
	TanggalDibuat     time.Time `db:"tanggal_dibuat" json:"tanggal_dibuat"`
}

func NewAgendaPertemuan(userId, dosenId int, judul, keterangan, tanggalPertemuan string) (*AgendaPertemuan, error) {
	parsedTanggalPertemuan, err := time.Parse("2006-01-02 15:04", tanggalPertemuan)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	formatedNow := now.Format("2006-01-02 15:04")
	tanggalDibuat, err := time.Parse("2006-01-02 15:04", formatedNow)
	if err != nil {
		return nil, err
	}

	agenda := &AgendaPertemuan{
		UserID:           userId,
		DosenID:          dosenId,
		Status:           "Menunggu persetujuan dosen",
		Judul:            judul,
		Keterangan:       keterangan,
		TanggalPertemuan: parsedTanggalPertemuan,
		TanggalDibuat:    tanggalDibuat,
	}

	if err := Dbm.Insert(agenda); err != nil {
		return nil, err
	}

	return agenda, nil
}

func AddCatatanAgendaPertemuan(agendaPertemuanId, userId int, catatan string) (*CatatanAgendaPertemuan, error) {
	now := time.Now()
	formatedNow := now.Format("2006-01-02 15:04")
	tanggalDibuat, err := time.Parse("2006-01-02 15:04", formatedNow)
	if err != nil {
		return nil, err
	}

	catatanAgendaPertemuan := &CatatanAgendaPertemuan{
		AgendaPertemuanID: agendaPertemuanId,
		UserID:            userId,
		Catatan:           catatan,
		TanggalDibuat:     tanggalDibuat,
	}

	if err := Dbm.Insert(catatanAgendaPertemuan); err != nil {
		return nil, err
	}

	return catatanAgendaPertemuan, nil

}
