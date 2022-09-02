package main

import (
	"fmt"
	"github.com/rxedu/go-bookstore/v1/pkg"
	"log"
	"net/http"
)

func main() {
	s := pkg.NewServer()
	pkg.InitDB()
	fmt.Printf("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", s))
}
