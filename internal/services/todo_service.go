package services

import (
	"errors"

	"github.com/kalanihaubrick/todo-api/internal/database"
	"github.com/kalanihaubrick/todo-api/internal/models"
)

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := database.DB.Find(&todos)
	return todos, result.Error
}

func GetTodoByID(id uint) (models.Todo, error) {
	var todo models.Todo
	result := database.DB.First(&todo, id)
	return todo, result.Error
}

func CreateTodo(todo models.Todo) error {
	var existingTodo models.Todo
	if err := database.DB.Where("task = ?", todo.Task).First(&existingTodo).Error; err == nil {
		return errors.New("a task with the same description already exists")
	}
	return database.DB.Create(&todo).Error
}

func UpdateTodo(todo models.Todo) error {
	var existingTodo models.Todo
	if err := database.DB.Where("task = ?", todo.Task).First(&existingTodo).Error; err == nil {
		return errors.New("a task with the same description already exists")
	}
	return database.DB.Save(&todo).Error
}

func DeleteTodo(id uint) error {
	var todo models.Todo
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		return result.Error
	}
	return database.DB.Delete(&todo).Error
}
