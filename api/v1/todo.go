package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kalanihaubrick/todo-api/internal/controllers"
	"github.com/kalanihaubrick/todo-api/internal/database"
	"github.com/kalanihaubrick/todo-api/internal/repositories"
	"github.com/kalanihaubrick/todo-api/internal/services"
)

func RegisterRoutes(r *gin.Engine) {
	db := database.DB
	repo := repositories.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	controller := controllers.NewTodoController(service)

	todoRoutes := r.Group("api/v1/todos")
	{
		todoRoutes.GET("", controller.GetAllTodos)
		todoRoutes.POST("", controller.CreateTodo)
		todoRoutes.GET("/:id", controller.GetTodoById)
		todoRoutes.PUT("/:id", controller.UpdateTodo)
		todoRoutes.DELETE("/:id", controller.DeleteTodo)
	}
}
