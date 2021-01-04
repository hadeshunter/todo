package database

import (
	"fmt"

	"github.com/hadeshunter/todo/models"
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
	// units := []models.Unit{}
	// for rows.Next() {
	// 	var unit models.Unit
	// 	if err := rows.Scan(&unit.DonviID, &unit.TenDV, &unit.DonviChaID, &unit.TenDVQL); err != nil {
	// 		return nil, err
	// 	}
	// 	units = append(units, unit)
	// }

	query := "SET SERVEROUTPUT ON SIZE 1000000 \n" +
		"DECLARE \n" +
		"o_cursor sys_refcursor; \n" +
		"o_donvi_id admin_hcm.donvi.donvi_id%TYPE; \n" +
		"o_ten_dv admin_hcm.donvi.ten_dv%TYPE; \n" +
		"BEGIN \n" +
		"khanhnv.dashboard.getTTVT(o_cursor); \n" +
		"LOOP \n" +
		"FETCH o_cursor \n" +
		"INTO o_donvi_id,o_ten_dv; \n" +
		"EXIT WHEN o_cursor%NOTFOUND; \n" +
		"END LOOP; \n" +
		"CLOSE o_cursor; \n" +
		"END; \n" +
		"/ \n"
	fmt.Printf(query)
	stmt, err := db.oracleDB.Prepare(query)

	// defer stmt.Close()
	result, err := stmt.Query(nil)
	if err != nil {
		fmt.Printf("Error execute query: %s\n", err)
		return nil, err
	}
	fmt.Println(result)
	return units, nil
}
