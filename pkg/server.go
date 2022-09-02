package pkg

import (
	"github.com/gorilla/mux"
	"github.com/rxedu/go-bookstore/v1/internal/models"
	"github.com/rxedu/go-bookstore/v1/internal/routes"
	"github.com/rxedu/go-bookstore/v1/pkg/config"
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
