package database

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"

	"github.com/hadeshunter/todo/models"
	go_ora "github.com/sijms/go-ora"

	// database driver
	_ "github.com/lib/pq"
)

// ListAllUnits ...
func (db *Database) ListAllUnits() ([]models.Unit, error) {
	units := []models.Unit{}

	conn, err := sql.Open("oracle", "oracle://trihm:2323@exax7-scan.vnpthcm.vn:1521/SGN")
	if err != nil {
		fmt.Println("Error connecting DB")
		return nil, err
	}
	defer conn.Close()

	// Create statment
	stmt, err := conn.Prepare("select donvi_id, ten_dv from ADMIN_HCM.donvi where donvi_id in (:1, :2, :3, :4, :5, :6, :7, :8, :9)")
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

	// Extract data using next
	for rows.Next() {
		var unit models.Unit
		if err := rows.Scan(&unit.DonviID, &unit.TenDV); err != nil {
			return nil, err
		}
		units = append(units, unit)
	}

	return units, nil
}

// ListAllTTVT ...
func (db *Database) ListAllTTVT() ([]models.Unit, error) {
	units := []models.Unit{}
	// Create statment
	cmdText := `BEGIN    
									test.getTTVT(:1, :2); 
							END;`
	stmt := go_ora.NewStmt(cmdText, db.oracleDB)
	stmt.AddParam("1", 41, 41, go_ora.Input)
	stmt.AddRefCursorParam("2")
	defer stmt.Close()

	// Query
	_, err := stmt.Exec(nil)
	if err != nil {
		fmt.Println("Error query")
		return nil, err
	}

	if cursor, ok := stmt.Pars[0].Value.(go_ora.RefCursor); ok {
		defer cursor.Close()
		rows, err := cursor.Query()
		if err != nil {
			return nil, err
		}

		var (
			donvi_id int64
			ten_dv   string
		)
		unit := models.Unit{}
		values := make([]driver.Value, 2)

		for {
			err = rows.Next(values)
			// check for error and if == io.EOF break
			if err == io.EOF {
				break
			}

			if donvi_id, ok = values[0].(int64); !ok {
				return nil, errors.New("Not have value DonviID")
			}

			if ten_dv, ok = values[1].(string); !ok {
				return nil, errors.New("Not have value TenDV")
			}

			unit.DonviID = donvi_id
			unit.TenDV = ten_dv
			units = append(units, unit)
		}
	}

	return units, nil
}
