package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/hadeshunter/todo/models"

	"github.com/jinzhu/gorm"
	// database driver
	_ "github.com/lib/pq"
)

// Database todo
type Database struct {
	instance *gorm.DB
	URL      string
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
		db.instance = postgresDB
	}
}

func (db *Database) migrate() {
	db.instance.AutoMigrate(&models.User{})
	db.instance.AutoMigrate(&models.Item{})
}

// ConnectOracle Connect Oracle Database
func ConnectOracle() (*sql.DB, error) {
	// Connect oracle database
	db, err := sql.Open("oci8", "khanhnv/2305@exax7-scan.vnpthcm.vn:1521/SGN")
	if err != nil {
		fmt.Println(err)
		return db, err
	}
	// defer db.Close()
	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return db, err
	}
	return db, err
}
