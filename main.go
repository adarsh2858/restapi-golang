package main

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID 		string `json:"id"`
	Isbn 	string `json:"isbn"`
	Title 	string `json:"title"`
	Author 	*Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// fmt.Println("Hello world")
	// Init router
	r := mux.NewRouter()

	// Mock data - @todo - implement db
	books = append(books, Book{ID: "1", Isbn: "448783", Title: "Book 1", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "341753", Title: "Book 2", Author: &Author {Firstname: "Jane", Lastname: "Salecha"}})

	// Route handlers / Endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
