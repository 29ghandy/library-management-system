package controllers

import (
	"awesomeProject/entities"
	"awesomeProject/services"
	"encoding/json"
	"net/http"
)

type BooksController struct{ fileName string }

func (c *BooksController) AddBook(response http.ResponseWriter, request *http.Request) {

	var book entities.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	var bookService services.BooksService
	bookService.AddBook(book)

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("book added"))
}
