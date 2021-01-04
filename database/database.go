package database

import (
	"fmt"
	"log"

	"github.com/hadeshunter/todo/models"
	"github.com/jinzhu/gorm"

	// database driver
	_ "github.com/lib/pq"
	go_ora "github.com/sijms/go-ora"
)

// Database todo
type Database struct {
	postgresDB *gorm.DB
	URL        string
	oracleDB   *go_ora.Connection
}

// New database
func New(url string) *Database {
	db := Database{URL: url}
	// db.initialize()
	// db.migrate()
	db.connectOracle()
	return &db
}

// ConnectOracle Connect Oracle Database
func (db *Database) connectOracle() {
	// Connect oracle database
	conn, err := go_ora.NewConnection("oracle://khanhnv:2305@exax7-scan.vnpthcm.vn:1521/SGN")
	// check for error
	err = conn.Open()
	if err != nil {
		fmt.Printf("Error connect to the database: %s\n", err)
		return
	}
	// defer conn.Close()
	db.oracleDB = conn
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
