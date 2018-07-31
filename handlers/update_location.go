package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

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

	lastUpdate := time.Now()
	formatedLastUpdate := lastUpdate.Format("2006-01-02 15:04")

	userID := r.Context().Value("user_id").(int)
	var status models.Status
	queryUser := "SELECT * FROM status WHERE user_id = ?"
	if err := models.Dbm.SelectOne(&status, queryUser, userID); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}
	parsedDate, err := time.Parse("2006-01-02 15:04", formatedLastUpdate)
	if err != nil {
		return
	}

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
	listAp := make(map[string]int)

	for _, apReq := range apRequests {
		for _, apData := range apDb {
			if apReq.BSSID == apData.BSSID {
				if value, ok := listAp[apReq.BSSID]; ok {
					value = value + apReq.Level
					listAp[apReq.BSSID] = value
				} else {
					listAp[apReq.BSSID] = apReq.Level
				}
			}
		}
	}

	if len(listAp) == 0 {
		// TODO UPDATE POSISI DOSEN TO ""
		status.Posisi = "Tidak ada"
		status.LastUpdate = parsedDate
		if _, err := models.Dbm.Update(&status); err != nil {
			errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
			return
		}

		errors.NewError("Tidak ada Access Point yang sesuai!", http.StatusBadRequest).WriteTo(w)
		return
	}

	// Deklarasi variabel untuk perhitungan data RGB
	var jarakRSSI float64 = 80 // batas jarak dari 0 - maksimum nilai level dbm yang diterima
	var minRSSI int = 0

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

	fmt.Println(len(listAp))

	for ap, dataLevel := range listAp {
		// Gedung A Lantai 1
		// A1 AP1 a0:3d:6f:85:c0:c1
		if ap == "a0:3d:6f:85:c0:c1" {
			average["A1AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// A1 AP2 a0:3d:6f:85:c3:01
		if ap == "a0:3d:6f:85:c3:01" {
			average["A1AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// A1 AP3 a0:e0:af:57:26:81
		if ap == "a0:e0:af:57:26:81" {
			average["A1AP3"] = Round((float64(dataLevel) / 10.0))
		}

		// Gedung A Lantai 2
		// A2 AP1 a0:3d:6f:8b:11:61
		if ap == "a0:3d:6f:8b:11:61" {
			average["A2AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// A2 AP2 a0:3d:6f:8f:96:01
		if ap == "a0:3d:6f:8f:96:01" {
			average["A2AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// A2 AP3 a0:3d:6f:8b:0f:21
		if ap == "a0:3d:6f:8b:0f:21" {
			average["A2AP3"] = Round((float64(dataLevel) / 10.0))
		}

		// Gedung E Lantai 1
		// E1 AP1 a0:3d:6f:8b:0b:c1
		if ap == "a0:3d:6f:8b:0b:c1" {
			average["E1AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// E1 AP2 a0:3d:6f:89:22:c1
		if ap == "a0:3d:6f:89:22:c1" {
			average["E1AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// E1 AP3 a0:3d:6f:89:22:e1
		if ap == "a0:3d:6f:89:22:e1" {
			average["E1AP3"] = Round((float64(dataLevel) / 10.0))
		}

		// Gedung E Lantai 2
		// E2 AP1 a0:3d:6f:8b:0d:81
		if ap == "a0:3d:6f:8b:0d:81" {
			average["E2AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// E2 AP2 a0:e0:af:9a:12:01
		if ap == "a0:e0:af:9a:12:01" {
			average["E2AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// E2 AP3 a0:e0:af:99:e4:e1
		if ap == "a0:e0:af:99:e4:e1" {
			average["E2AP3"] = Round((float64(dataLevel) / 10.0))
		}

		// F2 AP1 a0:3d:6f:8b:0d:c1
		if ap == "a0:3d:6f:8b:0d:c1" {
			average["F2AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// F2 AP2 a0:3d:6f:89:20:e1
		if ap == "a0:3d:6f:89:20:e1" {
			average["F2AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// F2 AP3 a0:3d:6f:8f:52:81
		if ap == "a0:3d:6f:8f:52:81" {
			average["F2AP3"] = Round((float64(dataLevel) / 10.0))
		}

		// F3 AP1 a0:3d:6f:85:c1:41
		if ap == "a0:3d:6f:85:c1:41" {
			average["F3AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// F3 AP2 a0:3d:6f:8f:92:21
		if ap == "a0:3d:6f:8f:92:21" {
			average["F3AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// F3 AP3 a0:3d:6f:5c:8a:a1
		if ap == "a0:3d:6f:5c:8a:a1" {
			average["F3AP3"] = Round((float64(dataLevel) / 10.0))
		}

		// F4 AP1 a0:3d:6f:8f:94:61
		if ap == "a0:3d:6f:8f:94:61" {
			average["F4AP1"] = Round((float64(dataLevel) / 10.0))
		}

		// F4 AP2 a0:3d:6f:76:b8:c1
		if ap == "a0:3d:6f:76:b8:c1" {
			average["F4AP2"] = Round((float64(dataLevel) / 10.0))
		}

		// F4 AP3 a0:3d:6f:85:c1:81
		if ap == "a0:3d:6f:85:c1:81" {
			average["F4AP3"] = Round((float64(dataLevel) / 10.0))
		}

	}

	fmt.Printf("average A1AP1: %d\n", average["A1AP1"])
	fmt.Printf("average A1AP2: %d\n", average["A1AP2"])
	fmt.Printf("average A1AP3: %d\n", average["A1AP3"])
	fmt.Printf("average A2AP1: %d\n", average["A2AP1"])
	fmt.Printf("average A2AP2: %d\n", average["A2AP2"])
	fmt.Printf("average A2AP3: %d\n", average["A2AP3"])
	fmt.Printf("average E1AP1: %d\n", average["E1AP1"])
	fmt.Printf("average E1AP2: %d\n", average["E1AP2"])
	fmt.Printf("average E1AP3: %d\n", average["E1AP3"])
	fmt.Printf("average E2AP1: %d\n", average["E2AP1"])
	fmt.Printf("average E2AP2: %d\n", average["E2AP2"])
	fmt.Printf("average E2AP3: %d\n", average["E2AP3"])
	fmt.Printf("average F2AP1: %d\n", average["F2AP1"])
	fmt.Printf("average F2AP2: %d\n", average["F2AP2"])
	fmt.Printf("average F2AP3: %d\n", average["F2AP3"])
	fmt.Printf("average F3AP1: %d\n", average["F3AP1"])
	fmt.Printf("average F3AP2: %d\n", average["F3AP2"])
	fmt.Printf("average F3AP3: %d\n", average["F3AP3"])
	fmt.Printf("average F4AP1: %d\n", average["F4AP1"])
	fmt.Printf("average F4AP2: %d\n", average["F4AP2"])
	fmt.Printf("average F4AP3: %d\n", average["F4AP3"])

	// Menentukan posisi gedung
	gedung := make(map[string]int)

	gedung["A1"] = average["A1AP1"] + average["A1AP2"] + average["A1AP3"]
	gedung["A2"] = average["A2AP1"] + average["A2AP2"] + average["A2AP3"]
	gedung["E1"] = average["E1AP1"] + average["E1AP2"] + average["E1AP3"]
	gedung["E2"] = average["E2AP1"] + average["E2AP2"] + average["E2AP3"]
	gedung["F2"] = average["F2AP1"] + average["F2AP2"] + average["F2AP3"]
	gedung["F3"] = average["F3AP1"] + average["F3AP2"] + average["F3AP3"]
	gedung["F4"] = average["F4AP1"] + average["F4AP2"] + average["F4AP3"]

	gedung["A"] = gedung["A1"] + gedung["A2"]
	gedung["E"] = gedung["E1"] + gedung["E2"]
	gedung["F"] = gedung["F2"] + gedung["F3"] + gedung["F4"]

	// check gedung
	if gedung["A"] < gedung["E"] && gedung["A"] < gedung["F"] {
		// AP Gedung A Lantai 1
		if gedung["A1"] < gedung["A2"] {
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["A1AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["A1AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["A1AP3"])
			lokasiGedung = "A1"
		} else if gedung["A2"] < gedung["A1"] { // lantai 2
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["A2AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["A2AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["A2AP3"])
			lokasiGedung = "A2"
		}
	} else if gedung["E"] < gedung["A"] && gedung["E"] < gedung["F"] {
		// AP Gedung E Lantai 1
		if gedung["E1"] < gedung["E2"] {
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["E1AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["E1AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["E1AP3"])
			lokasiGedung = "E1"
		} else if gedung["E2"] < gedung["E1"] { // lantai 2
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["E2AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["E2AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["E2AP3"])
			lokasiGedung = "E2"
		}
	} else if gedung["F"] < gedung["A"] && gedung["F"] < gedung["E"] {
		// AP Gedung F Lantai 2
		if gedung["F2"] < gedung["F3"] && gedung["F2"] < gedung["F4"] {
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["F2AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["F2AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["F2AP3"])
			lokasiGedung = "F2"
		} else if gedung["F3"] < gedung["F2"] && gedung["F3"] < gedung["F4"] { // lantai 3
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["F3AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["F3AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["F3AP3"])
			lokasiGedung = "F3"
		} else if gedung["F4"] < gedung["F2"] && gedung["F4"] < gedung["F3"] {
			normalisasi1 = int(jarakRSSI) + (minRSSI + average["F4AP1"])
			normalisasi2 = int(jarakRSSI) + (minRSSI + average["F4AP2"])
			normalisasi3 = int(jarakRSSI) + (minRSSI + average["F4AP3"])
			lokasiGedung = "F4"
		}
	}

	// Pelevelan persamaan (255/jarakRSSI)*normalisasi, nilai 255 didapatkan dari max nilai RGB
	level1 = Round((255 / jarakRSSI) * float64(normalisasi1)) // R
	level2 = Round((255 / jarakRSSI) * float64(normalisasi2)) // G
	level3 = Round((255 / jarakRSSI) * float64(normalisasi3)) // B

	// bandingkan data level RGB yang telah dihitung dengan data RGB Training di DB
	var dataRgb []models.DataRgb

	queryRgb := "SELECT * FROM data_rgb WHERE gedung=?"

	if _, err := models.Dbm.Select(&dataRgb, queryRgb, lokasiGedung); err != nil {
		errors.NewError("can't fetch rgb data", http.StatusInternalServerError).WriteTo(w)
		return
	}

	fmt.Printf("R: %d\n", level1)
	fmt.Printf("G: %d\n", level2)
	fmt.Printf("B: %d\n", level3)

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
	status.Posisi = posisi
	status.LastUpdate = parsedDate
	if _, err := models.Dbm.Update(&status); err != nil {
		errors.NewErrorWithStatusCode(http.StatusInternalServerError).WriteTo(w)
		return
	}
	var response updateResponse
	response.Posisi = posisi
	response.Mse = minimumValue

	json.NewEncoder(w).Encode(response)
}

// Round function to round float number to integer
func Round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
