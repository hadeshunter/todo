package main

import (
	"fmt"

	"github.com/hadeshunter/todo/database"
	"github.com/hadeshunter/todo/models"

	// database driver
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora"
)

func main() {
	// db := database.New(os.Getenv("DATABASE_URL"))

	// tasks := []string{
	// 	"Họp team training",
	// 	"Chạy thử todo",
	// 	"Viết code frontend",
	// 	"Viết code backend",
	// 	"Trả lời câu hỏi",
	// }

	// for index, task := range tasks {
	// 	item, err := db.CreateItem(task)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else if index < 3 {
	// 		db.ToggleItem(item.ID)
	// 	}
	// }

	// items, err := db.ListAllItems()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for index, item := range items {
	// 	if index+1 == 3 {
	// 		db.DeleteItem(item.ID)
	// 	}
	// }

	// litems, err := db.ListAllItems()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for index, item := range litems {
	// 	indicator := "☐"
	// 	if item.IsDone {
	// 		indicator = "☑︎"
	// 	}
	// 	fmt.Printf("%s %d. %s\n", indicator, index+1, item.Title)
	// }

	// godotenv.Load()
	// server := server.New()
	// server.Start(":5000")

	// Connect oracle database
	db, err := database.ConnectOracle()

	// Create statment
	stmt, err := db.Prepare("select donvi_id, ten_dv, donvi_cha_id, ten_dvql from ADMIN_HCM.donvi where donvi_id in (:1,:2,:3,:4,:5,:6,:7,:8,:9)")
	if err != nil {
		fmt.Println("Error create statment")
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	// Query
	rows, err := stmt.Query(41, 42, 43, 44, 45, 56, 57, 59, 60)
	if err != nil {
		fmt.Println("Error query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Extract data using next
	for rows.Next() {
		var unit models.Unit
		rows.Scan(&unit.DonviID, &unit.TenDV, &unit.DonviChaID, &unit.TenDVQL)
		println(unit.DonviID, unit.TenDV, unit.DonviChaID, unit.TenDVQL)
	}
}
