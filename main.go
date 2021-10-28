package main

import (
	"stephan/todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	uc := new(controllers.TodoController)

	router := gin.Default()
	router.GET("/todos", uc.GetAll)
	router.GET("todo/:id", uc.GetOne)
	router.POST("/todo", uc.AddOne)
	router.DELETE("/todo/:id", uc.DeleteOne)

	router.Run("localhost:8080")
}
