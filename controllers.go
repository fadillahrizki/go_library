package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getBooks(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	json.NewEncoder(response).Encode(Books)
}

func getBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	for _, book := range Books {
		if book.Title == params["title"] {
			json.NewEncoder(response).Encode(book)
			return
		}
	}
}

func createBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var book *Book
	book.Title = request.FormValue("title")
	book.Author.Name = request.FormValue("name")
	book.Author.Age = request.FormValue("age")
	Books = append(Books, book)
	json.NewEncoder(response).Encode(Books)
}
