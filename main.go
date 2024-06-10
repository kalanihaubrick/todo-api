package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Todo struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB
var err error

func main() {
	dsn := "sqlserver://kalani:123456@LAPTOP-VFC28UJV:1433?database=TodoDB"

	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.POST("/todos", createTodo)
	r.GET("/todos", getTodos)
	r.PUT("/todos/:id", updateTodo)
	r.GET("/todo/:id", getTodo)
	r.DELETE("/todo/:id", deleteTodo)

	r.Run(":8080")
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingTodo Todo

	if err := db.Where("task = ?", todo.Task).First(&existingTodo).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A task with the same description already exists"})
		return
	}

	db.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func getTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func updateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo Todo

	if result := db.First(&todo, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingTodo Todo

	if err := db.Where("task = ?", todo.Task).First(&existingTodo).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A task with the same description already exists"})
		return
	}

	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func getTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo Todo

	if result := db.First(&todo, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo Todo

	if result := db.First(&todo, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted: " + todo.Task})
}
