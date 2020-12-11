package database

import (
	"database/sql"
	"fmt"
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
	oracleDB   *sql.DB
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
	oracleDB, err := sql.Open("oracle", "oracle://khanhnv:2305@exax7-scan.vnpthcm.vn:1521/SGN")
	if err != nil {
		fmt.Printf("Error connect to the database: %s\n", err)
		return
	}
	// defer oracleDB.Close()

	// Ping oracle database
	if err = oracleDB.Ping(); err != nil {
		fmt.Printf("Error ping to the database: %s\n", err)
		return
	}
	db.oracleDB = oracleDB
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
