package main

import (
	"bookstore/models"
	"fmt"
	"net/http"
)

func main() {
	models.InitDB("postgres", "postgres://robinson:pword@localhost/bookstore")
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":9000", nil)
}
