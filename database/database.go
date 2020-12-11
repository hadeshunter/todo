package database

import (
	"log"

	"github.com/hadeshunter/todo/models"
	"github.com/jinzhu/gorm"

	// database driver
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora"
)

// Database todo
type Database struct {
	postgresDB *gorm.DB
	URL        string
}

// New database
func New(url string) *Database {
	db := Database{URL: url}
	db.initialize()
	db.migrate()
	return &db
}

func (db *Database) initialize() {
	if postgresDB, err := gorm.Open("postgres", db.URL); err != nil {
		log.Fatal(err)
	} else {
		db.postgresDB = postgresDB
	}
}

func (db *Database) migrate() {
	db.postgresDB.AutoMigrate(&models.User{})
	db.postgresDB.AutoMigrate(&models.Item{})
}
