package handlers

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type req struct {
	Data []data `json:"data"`
}

type data struct {
	BSSID string `json:"bssid"`
	SSID  string `json:"ssid"`
	Level int    `json:"level"`
}

type updateResponse struct {
	Posisi string  `json:"posisi"`
	Mse    float64 `json:"mse"`
}

// UpdateDosenLocation represent a reuqest for update location from dosen
func UpdateDosenLocation(w http.ResponseWriter, r *http.Request) {

	var req req

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		errors.NewErrorWithStatusCode(http.StatusBadRequest).WriteTo(w)
		return
	}

	apRequests := req.Data

	var apDb []models.AccessPoint

	query := "SELECT * FROM access_point"

	if _, err := models.Dbm.Select(&apDb, query); err != nil {
		errors.NewError("can't fetch access points", http.StatusInternalServerError).WriteTo(w)
		return
	}

	// filter AP yang sama dengan AP yang ada di DB
	listAp := make(map[data]map[string]int)

	for _, apReq := range apRequests {
		for _, apData := range apDb {
			if apReq.BSSID == apData.BSSID {
				if _, ok := listAp[apReq]; ok {
					temp := listAp[apReq]
					newMap := map[string]int{
						"count": temp["count"] + 1,
						"level": temp["level"] + apReq.Level,
					}
					listAp[apReq] = newMap
				} else {
					newMap := map[string]int{
						"count": 1,
						"level": apReq.Level,
					}
					listAp[apReq] = newMap
				}
			}
		}
	}

	if len(listAp) == 0 {
		// TODO UPDATE POSISI DOSEN TO ""
		errors.NewError("Tidak ada Access Point yang sesuai!", http.StatusBadRequest).WriteTo(w)
		return
	}

	// Deklarasi variabel untuk perhitungan data RGB
	jarakRSSI := 80 // batas jarak dari 0 - maksimum nilai level dbm yang diterima
	minRSSI := 0

	// nilai normalisasi
	var normalisasi1 int
	var normalisasi2 int
	var normalisasi3 int

	// level RGB
	var level1 int
	var level2 int
	var level3 int

	// lokasi gedung
	var lokasiGedung string
	var posisi string

	// Hitung rata2 tiap AP
	average := make(map[string]int)

	for ap, data := range listAp {
		// Gedung A Lantai 1
		// A1 AP1 a0:3d:6f:85:c0:c1
		if ap.BSSID == "a0:3d:6f:85:c0:c1" {
			average["A1AP1"] = data["level"] / 10
		}

		// A1 AP2 a0:3d:6f:85:c3:01
		if ap.BSSID == "a0:3d:6f:85:c3:01" {
			average["AP1AP2"] = data["level"] / 10
		}

		// A1 AP3 a0:e0:af:57:26:81
		if ap.BSSID == "a0:e0:af:57:26:81" {
			average["AP1AP3"] = data["level"] / 10
		}

		// Gedung A Lantai 2
		// A2 AP1 a0:3d:6f:8b:11:61
		if ap.BSSID == "a0:3d:6f:8b:11:61" {
			average["A1AP1"] = data["level"] / 10
		}

		// A2 AP2 a0:3d:6f:8f:96:01
		if ap.BSSID == "a0:3d:6f:8f:96:01" {
			average["A1AP1"] = data["level"] / 10
		}

		// A2 AP3 a0:3d:6f:8b:0f:21
		if ap.BSSID == "a0:3d:6f:8b:0f:21" {
			average["A1AP1"] = data["level"] / 10
		}
	}

	// Menentukan posisi gedung
	gedung := make(map[string]int)

	gedung["A1"] = average["A1AP1"] + average["A1AP2"] + average["A1AP3"]
	gedung["A2"] = average["A2AP1"] + average["A2AP2"] + average["A2AP3"]

	// AP Gedung A Lantai 1
	if gedung["A1"] < gedung["A2"] {
		normalisasi1 = jarakRSSI + (minRSSI + average["A1AP1"])
		normalisasi2 = jarakRSSI + (minRSSI + average["A1AP2"])
		normalisasi3 = jarakRSSI + (minRSSI + average["A1AP3"])
		lokasiGedung = "A1"
	} else if gedung["A2"] < gedung["A1"] {
		normalisasi1 = jarakRSSI + (minRSSI + average["A2AP1"])
		normalisasi2 = jarakRSSI + (minRSSI + average["A2AP2"])
		normalisasi3 = jarakRSSI + (minRSSI + average["A2AP3"])
		lokasiGedung = "A2"
	}

	// Pelevelan persamaan (255/jarakRSSI)*normalisasi, nilai 255 didapatkan dari max nilai RGB
	level1 = (255 / jarakRSSI) * normalisasi1 // R
	level2 = (255 / jarakRSSI) * normalisasi2 // G
	level3 = (255 / jarakRSSI) * normalisasi3 // B

	// bandingkan data level RGB yang telah dihitung dengan data RGB Training di DB
	var dataRgb []models.DataRgb

	queryRgb := "SELECT * FROM data_rgb WHERE gedung=?"

	if _, err := models.Dbm.Select(&dataRgb, queryRgb, lokasiGedung); err != nil {
		errors.NewError("can't fetch rgb data", http.StatusInternalServerError).WriteTo(w)
		return
	}

	// hitung mean squared error
	var mseData []float64

	// difference = rgb training - rgb testing
	for _, rgb := range dataRgb {
		// Mean Squared Error
		difLevel1 := rgb.LevelR - level1
		difLevel2 := rgb.LevelG - level2
		difLevel3 := rgb.LevelB - level3
		mse := (math.Pow(float64(difLevel1), 2) + math.Pow(float64(difLevel2), 2) + math.Pow(float64(difLevel3), 2)) / 3
		mseData = append(mseData, mse)
	}

	if len(mseData) == 0 {
		errors.NewError("Data MSE kosong", http.StatusInternalServerError).WriteTo(w)
		return
	}

	// hitung nilai minimum
	var minimumValue = mseData[0]
	var idxMinimum = 0

	for i, value := range mseData {
		if value < minimumValue {
			minimumValue = value
			idxMinimum = i
		}
	}

	// posisi dosen dilihat dari data RGB
	posisi = dataRgb[idxMinimum].Ruang
	//TODO UPDATE POSISI DOSEN
	var response updateResponse
	response.Posisi = posisi
	response.Mse = minimumValue

	json.NewEncoder(w).Encode(response)
}
