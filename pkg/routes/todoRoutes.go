package routes

import (
	"github.com/attanabilrabbani/go-todo/pkg/controllers"
	"github.com/gorilla/mux"
)

func TodoRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todo/add", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todo/", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", controllers.GetTodoById).Methods("GET")
	r.HandleFunc("/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", controllers.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/", controllers.Homepage)

	return r
}
