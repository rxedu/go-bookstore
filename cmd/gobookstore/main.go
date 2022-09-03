package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rxedu/go-bookstore"
)

func main() {
	s := gobookstore.NewServer()
	gobookstore.InitDB()
	fmt.Printf("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", s))
}
