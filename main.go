package main

import (
	"github.com/hadeshunter/todo/server"
	"os"
	"fmt"
	"log"
	"github.com/hadeshunter/todo/database"
)

func main() {
	db := database.New(os.Getenv("DATABASE_URL"))

	tasks := []string{
		"Họp team training",
		"Chạy thử todo",
		"Viết code frontend",
		"Viết code backend",
		"Trả lời câu hỏi",
	}

	for index, task := range tasks {
		item, err := db.CreateItem(task)
		if err != nil {
			log.Fatal(err)
		} else if index < 3 {
			db.CompleteItem(item.ID)
		}
	}

	items, err := db.ListAllItems()
	if err != nil {
		log.Fatal(err)
	}
	for index, item := range items {
		indicator := "☐"
		if item.IsDone {
			indicator = "☑︎"
		}
		fmt.Printf("%s %d. %s\n", indicator, index+1, item.Title)
	}

	server:= server.New()
	server.Start(":5000")
}