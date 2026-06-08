package main

import (
	"awesomeProject/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var mainRouter = mux.NewRouter()
	mainRouter.PathPrefix("/books").Handler(handler.BooksRouter())
	err := http.ListenAndServe(":8080", mainRouter)
	if err != nil {
		fmt.Println(err)
	}
}
