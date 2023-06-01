package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	//External library
	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"ID"`
	Title  string  `json:"Title"`
	Author *Author `json:"Author"`
	Price  int     `json:"Price"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Server converts the go data structure to JSON while sending it to Client and vice versa
var books = []Book{
	{ID: "1", Title: "book One", Author: &Author{Firstname: "X", Lastname: "Y"}, Price: 1000},
	{ID: "2", Title: "book Two", Author: &Author{Firstname: "M", Lastname: "P"}, Price: 150},
	{ID: "3", Title: "book Three", Author: &Author{Firstname: "D", Lastname: "P"}, Price: 1500},
}

// GET ALL BOOKS
func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET Book by ID
func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// POST NEW BOOK
func createbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// PUT UPDATE
func updatebook(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the books, range
	//delete the book with the i.d that we have sent
	//add a new book - the book which we send in the body of position

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// DELETE
func deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] { //... -> So on books
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", createbook).Methods("POST")
	r.HandleFunc("/books/{id}", updatebook).Methods("PUT")
	r.HandleFunc("/books/{id}", deletebook).Methods("DELETE")

	fmt.Printf("Starting server at port 9090\n")
	log.Fatal(http.ListenAndServe(":9090", r))

}
