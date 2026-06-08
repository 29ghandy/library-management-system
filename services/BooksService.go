package services

import (
	"awesomeProject/database"
	"awesomeProject/entities"
	"encoding/json"
)

type BooksService struct {
}

func (b *BooksService) AddBook(book entities.Book) error {
	fileHandler := database.File{}
	filename := "fileBase"

	// 1. Read existing data
	data, err := fileHandler.ReadFile(filename)
	if err != nil {
		// Initialize empty object if file doesn't exist
		data = map[string]interface{}{}
	}

	// 2. Marshal and unmarshal into a map
	jsonBytes, _ := json.Marshal(data)
	var dataMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &dataMap)
	if err != nil {
		return err
	}

	// 3. Access or create "books" slice
	var booksSlice []interface{}
	if val, ok := dataMap["books"]; ok {
		booksSlice, _ = val.([]interface{})
	} else {
		booksSlice = []interface{}{}
	}

	// 4. Append new book
	booksSlice = append(booksSlice, book)
	dataMap["books"] = booksSlice

	// 5. Write back the updated object
	err = fileHandler.WriteToFile(filename, dataMap)
	if err != nil {
		return err
	}

	return nil
}
