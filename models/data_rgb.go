package models

// DataRgb object
type DataRgb struct {
	ID     int    `db:"id" json:"id"`
	Gedung string `db:"gedung" json:"gedung"`
	Ruang  string `db:"ruang" json:"ruang"`
	LevelR int    `db:"level_r" json:"level_r"`
	LevelG int    `db:"level_g" json:"level_g"`
	LevelB int    `db:"level_b" json:"level_b"`
	AP1    int    `db:"ap1" json:"ap1"`
	AP2    int    `db:"ap2" json:"ap2"`
	AP3    int    `db:"ap3" json:"ap3"`
}

// NewDataRgb insert new row
func NewDataRgb(gedung, ruang string, level_r, level_g, level_b, ap1, ap2, ap3 int) (*DataRgb, error) {
	data := &DataRgb{
		Gedung: gedung,
		Ruang:  ruang,
		LevelR: level_r,
		LevelG: level_g,
		LevelB: level_b,
		AP1:    ap1,
		AP2:    ap2,
		AP3:    ap3,
	}

	if err := Dbm.Insert(data); err != nil {
		return nil, err
	}

	return data, nil
}
