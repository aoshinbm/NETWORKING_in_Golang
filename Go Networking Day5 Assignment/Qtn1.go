/*
Q1) Write a Go program that creates a HTTP server for an e-commerce website,
which listens on a specific port and responds to incoming HTTP
requests by fetching a slice of product information and
returning it to the client in a JSON format.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type PRODUCT struct {
	PRODUCT_ID   int
	PRODUCT_NAME string
}

var product PRODUCT

func main() {

	products = []PRODUCT{
		{PRODUCT_ID: 11, PRODUCT_NAME: "Clips"},
		{PRODUCT_ID: 22, PRODUCT_NAME: "Pencil"},
	}
	http.HandleFunc("/ecomm", handleProdInfo)
	fmt.Println("Starting a server at portNo 8082")
	http.ListenAndServe(":8082", nil)
}
func handleInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Listening..")
	switch r.Method {
	case "GET":
		handleGetProductInfo(w, r)
	default:
		http.Error(w, "INVALID", http.StatusMethodNotAllowed)
	}
}
func handleGetProductInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(product)
}

func handleProdInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(product)
	}
	http.Error(w, "Not found", http.StatusMethodNotAllowed)
}

func handleBuk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle book item api was called")
	//converting string into int
	//ranging whether its present or not
	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	fmt.Println(r.URL.Path[len("/books/"):])
	path := "/books/12"
	bookid := path[8:]
	fmt.Println(bookid)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, bookk := range books {
		if bookk.BOOK_ID == id {
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(bookk)
			case "PUT":
				var newBook BOOK
				json.NewDecoder(r.Body).Decode(&newBook)
				newBook.BOOK_ID = id
				books[i] = newBook
				json.NewEncoder(w).Encode(newBook)
			case "DELETE":
				books = append(books[:i], books...)
				w.WriteHeader(http.StatusNoContent)
			default:
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)

}
