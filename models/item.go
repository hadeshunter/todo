package models

import (
	"github.com/jinzhu/gorm"
)

// Item of list
type Item struct {
	gorm.Model
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
