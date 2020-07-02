package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Company contains information to export VAT invoice
type Company struct {
	gorm.Model
	TaxNumber string `json:"tax_number"`
	Name      string `json:"company_name"`
	Address   string `json:"company_address"`
	IsDefault bool   `json:"is_default"`
	UserID    uint   `json:"-"`
}

// User contains all information of our User
type User struct {
	gorm.Model
	Name      string
	Email     string
	Phone     string
	Companies []Company `json:"companies" gorm:"foreignkey:UserID"`
}

// MarshalJSON return the json object of user
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct{
		ID			uint			`json:"id"`
		Name		string		`json:"name"`
		Phone		string		`json:"phone"`
		Email		string		`json:"email"`
	}{
		ID:			u.ID,
		Name:		u.Name,
		Phone:	u.Phone,
		Email:	u.Email,
	})
}