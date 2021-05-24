package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	// "math/rand"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "github.com/aws/aws-lambda-go/lambda"
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
	// r.HandleFunc("/api/books", GetBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	// r.HandleFunc("/api/books", CreateBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")

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

func connectToDB() {
	db, err := sql.Open("mysql", "root:insert_password@tcp(127.0.0.1:3306)/go_tutorial")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	showTables, err := db.Query("SHOW TABLES")

	if err != nil {
		panic(err.Error())
	}

	for showTables.Next() {
		var tableName string

		err = showTables.Scan(&tableName)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(tableName)
	}

	// insert, err := db.Query("INSERT INTO foobar VALUES('LIPSUM2')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	results, err := db.Query("SELECT * FROM foobar")

	for results.Next() {
		var user string
		err = results.Scan(&user)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user)
	}

	fmt.Println("Successfully fetch the element")
}

func calc(x, y float64) (sum, difference float64) {
	sum = x + y
	difference = math.Abs(x - y)
	return
}

func main() {
	// RESTAPI implementation calls
	// InitialMigration()
	// Init router
	// InitializeRouter()
	// PracticeDeferWithCopyFile("go-file-1.txt", "go-file-2.txt")
	// lambda.Start(Handler)
	// connectToDB()
	res1, res2 := calc(1.1, 2.2)
	fmt.Printf("%.2f %.2f", res1, res2)
}
