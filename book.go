package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:@tcp(127.0.0.1:3306)/godb?parseTime=true"

// Book struct (Model)
type Book struct {
	gorm.Model
	// ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Book{})
}

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var books []Book
	DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// Get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	params := mux.Vars(r) // Get params
	var book Book
	DB.First(&book, params["id"])
	json.NewEncoder(w).Encode(book)
	// Loop through books and find with id
	// for _, item := range books {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	DB.Create(&book)
	// book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock Id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// for index, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:index], books[index+1:]...)
	// 		var book Book
	// 		_ = json.NewDecoder(r.Body).Decode(&book)
	// 		book.ID = params["id"]
	// 		books = append(books, book)
	// 		json.NewEncoder(w).Encode(book)
	// 		return
	// 	}
	// }
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// for index, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:index], books[index+1:]...)
	// 		break
	// 	}
	// }
	json.NewEncoder(w).Encode(books)
}
