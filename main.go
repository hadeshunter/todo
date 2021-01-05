package main

import (
	"log"
	"os"

	"github.com/hadeshunter/todo/database"

	// database driver
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora"
)

func main() {
	db := database.New(os.Getenv("DATABASE_URL"))
	_, err := db.ListAllUnits()
	if err != nil {
		log.Fatal(err)
	}

	// for _, unit := range listUnits {
	// 	fmt.Printf("%d | %s\n", unit.DonviID, unit.TenDV)
	// }
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

	// listItems, err := db.ListAllItems()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for index, item := range listItems {
	// 	indicator := "☐"
	// 	if item.IsDone {
	// 		indicator = "☑︎"
	// 	}
	// 	fmt.Printf("%s %d. %s\n", indicator, index+1, item.Title)
	// }

	// godotenv.Load()
	// server := server.New()
	// server.Start(":5000")
}
