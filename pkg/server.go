package pkg

import (
	"github.com/gorilla/mux"
	"github.com/rxedu/go-bookstore/v1/internal/routes"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	return r
}
