package database

import (
	"database/sql"
	"fmt"

	// database driver
	_ "github.com/sijms/go-ora"
)

// ConnectOracle Connect Oracle Database
func ConnectOracle() (*sql.DB, error) {
	// Connect oracle database
	db, err := sql.Open("oracle", "oracle://khanhnv:2305@exax7-scan.vnpthcm.vn:1521/SGN")
	if err != nil {
		fmt.Printf("Error connect to the database: %s\n", err)
		return nil, err
	}
	defer db.Close()

	// Ping oracle database
	if err = db.Ping(); err != nil {
		fmt.Printf("Error ping to the database: %s\n", err)
		return nil, err
	}
	return db, nil
}
