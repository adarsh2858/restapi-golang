package main

import (
	// "fmt"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// "math/rand"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

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

func PracticeDeferWithCopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// Calculate function below is to run a simple test
func Calculate(num int) int {
	return num + 2
}

func Add(x, y int) int {
	return x + y
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	// RESTAPI implementation calls
	// InitialMigration()
	// Init router
	// InitializeRouter()
	// PracticeDeferWithCopyFile("go-file-1.txt", "go-file-2.txt")
	lambda.Start(Handler)
}
