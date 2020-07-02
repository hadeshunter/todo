package database

import (
	"github.com/hadeshunter/todo/models"
)

// CreateItem is add new item
func (db *Database) CreateItem(title string) (*models.Item, error) {
	item := models.Item{
		Title:  title,
		IsDone: false,
	}
	if err := db.instance.Where(&models.Item{Title: title}).FirstOrCreate(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// DeleteItem delete item from list
func (db *Database) DeleteItem(id uint) (*models.Item, error) {
	var item models.Item
	if err := db.instance.Where("id = ?", id).Delete(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// ToggleItem change current status of item
func (db *Database) ToggleItem(id uint) (*models.Item, error) {
	var item models.Item
	if err := db.instance.First(&item, id).Error; err != nil {
		return nil, err
	}
	item.IsDone = !item.IsDone
	if err := db.instance.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// CompleteItem is set stastus of item to true when check
func (db *Database) CompleteItem(id uint) (*models.Item, error) {
	var item models.Item
	if err := db.instance.First(&item, id).Error; err != nil {
		return nil, err
	}
	item.IsDone = true
	if err := db.instance.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// ListAllItems is the function that get all item
func (db *Database) ListAllItems() ([]models.Item, error) {
	var items []models.Item
	if err := db.instance.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
