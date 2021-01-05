package database

import (
	"fmt"

	"github.com/hadeshunter/todo/models"
	go_ora "github.com/sijms/go-ora"

	// database driver
	_ "github.com/lib/pq"
)

// ListAllUnits ...
func (db *Database) ListAllUnits() ([]models.Unit, error) {
	units := []models.Unit{}
	// // Create statment
	// stmt, err := db.oracleDB.Prepare("select donvi_id, ten_dv, donvi_cha_id, ten_dvql from ADMIN_HCM.donvi where donvi_id in (:1, :2, :3, :4, :5, :6, :7, :8, :9)")
	// if err != nil {
	// 	fmt.Println("Error create statment")
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// defer stmt.Close()

	// // Query
	// rows, err := stmt.Query(41, 42, 43, 44, 45, 56, 57, 59, 60)
	// if err != nil {
	// 	fmt.Println("Error query")
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// defer rows.Close()

	// // Extract data using next
	// for rows.Next() {
	// 	var unit models.Unit
	// 	if err := rows.Scan(&unit.DonviID, &unit.TenDV, &unit.DonviChaID, &unit.TenDVQL); err != nil {
	// 		return nil, err
	// 	}
	// 	units = append(units, unit)
	// }

	cmdText := `
	DECLARE
		o_cursor sys_refcursor;
	BEGIN
		dashboard.getTTVT(o_cursor);
		LOOP
			FETCH o_cursor 
			INTO :donvi_id,:ten_dv;
			EXIT WHEN o_cursor%NOTFOUND;
    END LOOP;
    CLOSE o_cursor;
	END;`

	stmt := go_ora.NewStmt(cmdText, db.oracleDB)
	// defer stmt.Close()
	stmt.AddParam("donvi_id", "", 1000, go_ora.Output)
	stmt.AddParam("ten_dv", "", 1000, go_ora.Output)
	_, err := stmt.Exec(nil)
	if err != nil {
		return nil, err
	}

	for _, par := range stmt.Pars {
		fmt.Println(par.Value)
	}
	return units, nil
}
