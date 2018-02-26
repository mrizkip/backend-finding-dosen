package main

import (
	"time"

	"github.com/mrizkip/backend-finding-dosen/models"
)

func main() {
	if err := models.CreateTables(); err != nil {
		panic(err)
	}

	models.NewUser("test@mail.com", "123321", "Test Account", "NIM", "145110101111111", "081234567890", "mahasiswa") // 1
	models.NewUser("test1@mail.com", "123321", "Test Account 1", "NIM", "145110101111111", "081234567891", "mahasiswa")

	models.NewUser("dosen@mail.com", "123321", "Test Account Dosen", "NIP", "8912830812387192837", "081234567892", "dosen") // 2
	models.NewUser("dosen1@mail.com", "123321", "Test Account Dosen 1", "NIP", "8912830812387192838", "081234567893", "dosen")
	models.NewUser("dosen2@mail.com", "123321", "Test Account Dosen 2", "NIP", "8912830812387192839", "081234567894", "dosen")
	models.NewUser("dosen3@mail.com", "123321", "Test Account Dosen 3", "NIP", "8912830812387192840", "081234567895", "dosen")
	models.NewUser("dosen4@mail.com", "123321", "Test Account Dosen 4", "NIP", "8912830812387192841", "081234567896", "dosen")

	lastUpdate := time.Now()
	formatedLastUpdate := lastUpdate.Format("2006-01-02 15:04")
	models.NewStatus(3, "Tidak Aktif", "", formatedLastUpdate)
	models.NewStatus(4, "Tidak Aktif", "", formatedLastUpdate)
	models.NewStatus(5, "Tidak Aktif", "", formatedLastUpdate)
	models.NewStatus(6, "Tidak Aktif", "", formatedLastUpdate)
	models.NewStatus(7, "Tidak Aktif", "", formatedLastUpdate)
}
