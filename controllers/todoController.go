package controllers

import (
	"net/http"
	"stephan/todo/models"

	"github.com/gin-gonic/gin"
)

var todoModel = new(models.Todo)

type TodoController struct{}

func (t TodoController) GetAll(c *gin.Context) {
	todos := todoModel.GetAllTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func (t TodoController) GetOne(c *gin.Context) {
	todo := todoModel.GetByID(c.Param("id"))
	c.IndentedJSON(http.StatusOK, todo)
}

func (t TodoController) AddOne(c *gin.Context) {
	todo := todoModel.InsertOne(c)
	c.IndentedJSON(http.StatusCreated, todo)
}

func (t TodoController) DeleteOne(c *gin.Context) {
	_, err := todoModel.DeleteOne(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted one document"})
}
