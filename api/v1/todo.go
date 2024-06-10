package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kalanihaubrick/todo-api/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/todos", controllers.GetTodos)
		v1.GET("/todo/:id", controllers.GetTodoById)
		v1.POST("/todos", controllers.CreateTodo)
		v1.PUT("/todo/:id", controllers.UpdateTodo)
		v1.DELETE("/todo/:id", controllers.DeleteTodo)
	}
}
