package main

import (
	"log"
	"net/http"

	"github.com/attanabilrabbani/go-todo/pkg/config"
	"github.com/attanabilrabbani/go-todo/pkg/middleware"
	"github.com/attanabilrabbani/go-todo/pkg/models"
	"github.com/attanabilrabbani/go-todo/pkg/routes"
)

func main() {
	config.DBConnect()
	db := config.GetDB()
	db.AutoMigrate(&models.Todo{})

	r := routes.TodoRoutes()

	corsRouter := middleware.EnableCORSfunc(r)

	log.Println("Starting server on port:8000")
	log.Fatal(http.ListenAndServe(":8000", corsRouter))
}
