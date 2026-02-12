package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func main() {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return
	}

	var record Record

	err = json.Unmarshal(fileData, &record)
	if err != nil {
		fmt.Println("Ошибка парсинга Json", err)
		return
	}

	fmt.Printf("ID: %d, TITLE: %s", record.Id, record.Title)
}