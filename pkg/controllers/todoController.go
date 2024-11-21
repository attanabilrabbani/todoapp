package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/attanabilrabbani/go-todo/pkg/config"
	"github.com/attanabilrabbani/go-todo/pkg/models"
	"github.com/gorilla/mux"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err2 := models.CreateTodo(config.GetDB(), &todo)
	if err2 != nil {
		http.Error(w, "Failed to add new todo", http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"data": todo,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAllTodo(config.GetDB())
	if err != nil {
		http.Error(w, "Failed to fetch todo information", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)

}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	Id, err := strconv.Atoi(todoId)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	todoDetails, err := models.GetTodoById(config.GetDB(), Id)
	if err != nil {
		http.Error(w, "Error: Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoDetails)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	Id, err := strconv.Atoi(todoId)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	//decode the updated Todo
	var todoUpdate models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todoUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoUpdate.ID = uint(Id)

	if err := models.UpdateTodo(config.GetDB(), &todoUpdate); err != nil {
		http.Error(w, "Todo Update Failed", http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"data": todoUpdate,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	Id, err := strconv.Atoi(todoId)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err2 := models.DeleteTodo(config.GetDB(), Id)
	if err2 != nil {
		http.Error(w, "Todo Delete Failed", http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"message": "todo deleted",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
