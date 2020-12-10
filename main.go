package main

import (
	"fmt"

	"github.com/hadeshunter/todo/database"
	_ "github.com/mattn/go-oci8"
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
	if err != nil {
		return
	}
	rows, err := db.Query("select donvi_id, ten_dv, donvi_cha_id, ten_dvql from ADMIN_HCM.donvi where donvi_cha_id = :1 and donvi_ql = :2", 41, 0)
	if err != nil {
		fmt.Println("Error fetching ADMIN_HCM.donvi")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var donviID int
		var tenDV string
		var donviChaID int
		var tenDVQL string
		rows.Scan(&donviID, &tenDV, &donviChaID, &tenDVQL)
		println(donviID, tenDV, donviChaID, tenDVQL)
	}
}
