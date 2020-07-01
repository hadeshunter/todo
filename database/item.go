package database

import (
	"github.com/hadeshunter/todo/models"
)
// CreateItem is add new item
func (db *Database) CreateItem(title string) (*models.Item, error) {
	item := models.Item{
		Title: title,
		IsDone: false,
	}
	if err := db.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
// CompleteItem is set stastus of item to true when check
func (db *Database) CompleteItem(id uint) error {
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		return err
	}
	item.IsDone = true
	if err := db.Save(&item).Error; err != nil {
		return err
	}
	return nil
}
// ListAllItems is the function that get all item
func (db *Database) ListAllItems() ([]models.Item, error) {
	var items []models.Item
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}