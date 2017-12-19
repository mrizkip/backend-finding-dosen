package main

import (
	"github.com/mrizkip/backend-finding-dosen/models"
)

func main() {
	if err := models.CreateTables(); err != nil {
		panic(err)
	}

	models.NewUser("test@mail.com", "123321", "Test Account", "NIM", "145110101111111", "mahasiswa") // 1
	models.NewUser("test1@mail.com", "123321", "Test Account 1", "NIM", "145110101111111", "mahasiswa")

	models.NewUser("dosen@mail.com", "123321", "Test Account Dosen", "NIP", "8912830812387192837", "dosen") // 2
	models.NewUser("dosen1@mail.com", "123321", "Test Account Dosen 1", "NIP", "8912830812387192838", "dosen")
	models.NewUser("dosen2@mail.com", "123321", "Test Account Dosen 2", "NIP", "8912830812387192839", "dosen")
	models.NewUser("dosen3@mail.com", "123321", "Test Account Dosen 3", "NIP", "8912830812387192840", "dosen")
	models.NewUser("dosen4@mail.com", "123321", "Test Account Dosen 4", "NIP", "8912830812387192841", "dosen")
}
