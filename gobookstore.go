package gobookstore

import (
	"github.com/gorilla/mux"

	"github.com/rxedu/go-bookstore/internal/config"
	"github.com/rxedu/go-bookstore/internal/models"
	"github.com/rxedu/go-bookstore/internal/routes"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	return r
}

func InitDB() {
	config.ConnectDB()
	db := config.GetDB()
	models.InitModels(db)
}
