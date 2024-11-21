package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"size:50;default:'pending'" json:"status"`
	Priority    string    `gorm:"size:50;default:'normal'" json:"priority"`
	DueDate     time.Time `json:"due_date"`
}

func GetAllTodo(db *gorm.DB) ([]Todo, error) {
	var todos []Todo
	err := db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodoById(db *gorm.DB, Id int) (*Todo, error) {
	var getTodoId Todo
	err := db.Where("ID=?", Id).Find(&getTodoId).Error
	if err != nil {
		return nil, err
	}
	return &getTodoId, err
}

func CreateTodo(db *gorm.DB, todo *Todo) error {
	todo.DueDate = todo.DueDate.UTC()
	return db.Create(todo).Error
}

func UpdateTodo(db *gorm.DB, todo *Todo) error {
	var existingTodo Todo
	//find if db exists
	if err := db.First(&existingTodo, "ID = ?", todo.ID).Error; err != nil {
		return err
	}

	updateData := make(map[string]interface{})

	if todo.Title != "" {
		updateData["title"] = todo.Title
	}
	if todo.Description != "" {
		updateData["description"] = todo.Description
	}
	if todo.Status != "" {
		updateData["status"] = todo.Status
	}
	if todo.Priority != "" {
		updateData["priority"] = todo.Priority
	}
	if !todo.DueDate.IsZero() { // Check if the DueDate is non zero
		updateData["due_date"] = todo.DueDate
	}

	if len(updateData) > 0 {
		err := db.Model(&existingTodo).Updates(updateData).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteTodo(db *gorm.DB, Id int) error {
	var todos Todo
	return db.Where("ID=?", Id).Delete(todos).Error

}
