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

var records []Record

func main() {

	readFile(&records)

	
	// fmt.Printf("ID: %d, TITLE: %s", record.Id, record.Title)
	fmt.Printf("Your records:\n")
	fmt.Printf("__________________________________________________________________________________________\n")
	fmt.Printf("\n")
	data, _ := json.Marshal(records)
	fmt.Println(string(data))
	fmt.Printf("__________________________________________________________________________________________\n")
	fmt.Printf("\n")
	

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

			createRecord(&records, titleNewRecord, descriptionNewRecord)
		}
	}
}

func saveFile(records []Record)  {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Printf("Ошибка создания файла", err)
	}
	defer file.Close()

	data, _ := json.Marshal(records)
	os.WriteFile("data.json", data, 0644)
}

func createRecord(record *[]Record,title string, description string)  {
	newRecord := Record{
		Title: title,
		Description: description,
	}

	*record = append(*record, newRecord)
	saveFile(records)
}

func readFile(record *[]Record)  {
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