package main

import (
	// "fmt"
	"log"
	"net/http"

	// "math/rand"
	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	// Mock data - @todo - implement db
	// books = append(books, Book{ID: "1", Isbn: "448783", Title: "Book 1"})
	// books = append(books, Book{ID: "2", Isbn: "341753", Title: "Book 2"})
	// books = append(books, Book{ID: "1", Isbn: "448783", Title: "Book 1", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "341753", Title: "Book 2", Author: &Author{Firstname: "Jane", Lastname: "Salecha"}})

	// Route handlers / Endpoint
	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/books", CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	// fmt.Println("Hello world")
	InitialMigration()
	// Init router
	InitializeRouter()
}
