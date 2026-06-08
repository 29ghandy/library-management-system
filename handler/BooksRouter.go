package handler

import (
	"awesomeProject/controllers"

	"github.com/gorilla/mux"
)

func BooksRouter() *mux.Router {
	controller := controllers.BooksController{}
	router := mux.NewRouter().PathPrefix("/books").Subrouter()
	router.HandleFunc("/add-book", controller.AddBook).Methods("POST")
	return router
}
