package models

import (
	"encoding/json"
)

// Unit of list
type Unit struct {
	DonviID int
	TenDV   string
}

// MarshalJSON return the json object of Unit
func (u *Unit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		DonviID int    `json:"donvi_id"`
		TenDV   string `json:"ten_dv"`
	}{
		DonviID: u.DonviID,
		TenDV:   u.TenDV,
	})
}
