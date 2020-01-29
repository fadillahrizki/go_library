package main

import (
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func router() {
	Router.HandleFunc("/", getBooks).Methods("GET")
	Router.HandleFunc("/getBook/{title}", getBook).Methods("GET")
	Router.HandleFunc("/createBook", createBook).Methods("POST")
}
