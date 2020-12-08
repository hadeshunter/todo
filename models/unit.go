package models

import (
	"encoding/json"
)

// Unit of list
type Unit struct {
	DonviID    int
	TenDV      string
	DonviChaID int
	TenDVQL    string
}

// MarshalJSON return the json object of Unit
func (u *Unit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		DonviID    int    `json:"donvi_id"`
		TenDV      string `json:"ten_dv"`
		DonviChaID int    `json:"donvi_cha_id"`
		TenDVQL    string `json:"ten_dvql"`
	}{
		DonviID:    u.DonviID,
		TenDV:      u.TenDV,
		DonviChaID: u.DonviChaID,
		TenDVQL:    u.TenDVQL,
	})
}
