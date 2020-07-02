package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Item of list
type Item struct {
	gorm.Model
	Title  string
	IsDone bool
}

// MarshalJSON return the json object of item
func (i *Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID			uint		`json:"id"`
		Title		string	`json:"name"`
		IsDone	bool		`json:"is_done"`
	}{
		ID:			i.ID,
		Title:	i.Title,
		IsDone:	i.IsDone,
	})
}
