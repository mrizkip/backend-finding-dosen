package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrizkip/backend-finding-dosen/errors"
	"github.com/mrizkip/backend-finding-dosen/models"
)

type req struct {
	Data []data `json:"data"`
}

type data struct {
	BSSID     string `json:"bssid"`
	SSID      string `json:"ssid"`
	Frequency string `json:"frequency"`
	Level     int    `json:"level"`
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

}
