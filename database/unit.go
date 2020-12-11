package database

import (
	"fmt"

	"github.com/hadeshunter/todo/models"
)

// ListAllUnits ...
func (db *Database) ListAllUnits() ([]models.Unit, error) {
	// Create statment
	stmt, err := db.oracleDB.Prepare("select donvi_id, ten_dv, donvi_cha_id, ten_dvql from ADMIN_HCM.donvi where donvi_id in (:1,:2,:3,:4,:5,:6,:7,:8,:9)")
	if err != nil {
		fmt.Println("Error create statment")
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	// Query
	rows, err := stmt.Query(41, 42, 43, 44, 45, 56, 57, 59, 60)
	if err != nil {
		fmt.Println("Error query")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	units := []models.Unit{}
	// Extract data using next
	for rows.Next() {
		var unit models.Unit
		if err := rows.Scan(&unit.DonviID, &unit.TenDV, &unit.DonviChaID, &unit.TenDVQL); err != nil {
			return nil, err
		}
		units = append(units, unit)
	}
	return units, nil
}
