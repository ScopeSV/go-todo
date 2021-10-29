package main

import (
	"stephan/todo/controllers"
	"stephan/todo/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	uc := new(controllers.TodoController)

	router := gin.Default()
	router.Use(middlewares.IsAuthenticated)
	router.GET("/todos", uc.GetAll)
	router.GET("todo/:id", uc.GetOne)
	router.POST("/todo", uc.AddOne)
	router.DELETE("/todo/:id", uc.DeleteOne)
	router.PATCH("/todo/:id/complete", uc.ToggleComplete)

	router.Run("localhost:8080")
}
