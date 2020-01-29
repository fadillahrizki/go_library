package main

import (
	"fmt"
	"net/http"
)

func main() {
	router()
	fmt.Println("Starting server...")
	http.ListenAndServe(":8000", Router)
}
