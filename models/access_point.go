package models

// AccessPoint represent Access Point object
type AccessPoint struct {
	ID        int    `db:"id" json:"id"`
	BSSID     string `db:"bssid" json:"bssid"`
	SSID      string `db:"ssid" json:"ssid"`
	Frequency string `db:"frequency" json:"frequency"`
}

// NewAccessPoint is function for add new row in AccessPoint database table
func NewAccessPoint(bssid, ssid, frequency string) (*AccessPoint, error) {
	ap := &AccessPoint{
		BSSID:     bssid,
		SSID:      ssid,
		Frequency: frequency,
	}

	if err := Dbm.Insert(ap); err != nil {
		return nil, err
	}

	return ap, nil
}
