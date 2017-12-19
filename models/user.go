package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int    `db:"id" json:"id"`
	Email          string `db:"email" json:"email"`
	Password       string `db:"password" json:"password"`
	Nama           string `db:"nama" json:"nama"`
	JenisIdentitas string `db:"jenis_identitas" json:"jenis_identitas"`
	NoIdentitas    string `db:"no_identitas" json:"no_identitas"`
	Role           string `db:"role" json:"role"`
}

func NewUser(email, password, nama, jenisIdentitas, noIdentitas, role string) (*User, error) {
	user := &User{
		Email:          email,
		Nama:           nama,
		JenisIdentitas: jenisIdentitas,
		NoIdentitas:    noIdentitas,
		Role:           role,
	}
	user.HashPassword(password)
	err := Dbm.Insert(user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *User) HashPassword(raw string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hashedPassword)
}

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
