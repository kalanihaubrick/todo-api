package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kalanihaubrick/todo-api/internal/models"

	"github.com/kalanihaubrick/todo-api/internal/services"

	"net/http"
	"strconv"
)

func GetTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := services.GetTodoByID(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateTodo(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := services.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
