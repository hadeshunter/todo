package main

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/hadeshunter/todo/database"
	"github.com/hadeshunter/todo/server"
)

func main() {
	db := database.New(os.Getenv("DATABASE_URL"))

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

	items, err := db.ListAllItems()
	if err != nil {
		log.Fatal(err)
	}

	for index, item := range items {
		if index+1 == 3 {
			db.DeleteItem(item.ID)
		}
	}

	litems, err := db.ListAllItems()
	if err != nil {
		log.Fatal(err)
	}

	for index, item := range litems {
		indicator := "☐"
		if item.IsDone {
			indicator = "☑︎"
		}
		fmt.Printf("%s %d. %s\n", indicator, index+1, item.Title)
	}

	godotenv.Load()
	server:= server.New()
	server.Start(":5000")
}