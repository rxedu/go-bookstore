package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rxedu/go-bookstore/v1/pkg"
)

func main() {
	s := pkg.NewServer()
	pkg.InitDB()
	fmt.Printf("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", s))
}
