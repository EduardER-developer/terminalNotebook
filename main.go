package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Title 		string 	`json:"title"`
	Description string 	`json:"description"`
}

func main() {

	var record Record
	readFile(&record)

	
	// fmt.Printf("ID: %d, TITLE: %s", record.Id, record.Title)

	data, _ := json.Marshal(record)
	fmt.Println(string(data))
	

	for {
		fmt.Printf("Send 0-create new notebook, else send title record\n")
		selectRecord := ""
		fmt.Scan(&selectRecord)

		if selectRecord == "0" {
			fmt.Printf("You select command - new notebook\n")
			fmt.Printf("Enter you record title:\n")
			titleNewRecord := ""
			fmt.Scan(&titleNewRecord)
			fmt.Printf("Enter you record description:\n")
			descriptionNewRecord := ""
			fmt.Scan(&descriptionNewRecord)

			createRecord()
		}
	}
}

func createRecord()  {
	
}

func readFile(record *Record)  {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return
	}

	err = json.Unmarshal(fileData, record)
	if err != nil {
		fmt.Println("Ошибка парсинга Json", err)
		return
	}
}