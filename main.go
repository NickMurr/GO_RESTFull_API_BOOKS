package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go_course/2-rest-api/book-list/controllers"
	"github.com/go_course/2-rest-api/book-list/driver"
	"github.com/go_course/2-rest-api/book-list/models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

// Book structure

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()

	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
