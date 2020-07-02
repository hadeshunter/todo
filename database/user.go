package database

import (
	"github.com/hadeshunter/todo/models"
)

// GetUserByPhone ..
func (db *Database) GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	if err := db.instance.
		Where(&models.User{Phone: phone}).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser ..
func (db *Database) CreateUser(name string, phone string, email string) (*models.User, error) {
	user := models.User{
		Name: name,
		Phone: phone,
		Email: email,
	}
	if err := db.instance.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers ..
func (db *Database) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.instance.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}