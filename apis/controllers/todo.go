package controllers

import (
	"apis/db"
	"apis/models"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AddTodoController, handler for todo post api
func AddTodoController(c *gin.Context) {
	context := context.Background()
	todo := &models.Todo{}

	if err := c.BindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if strings.TrimSpace(todo.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Task title required",
		})

		return
	}

	userID, _ := c.Get("user_id")
	todo.UserID = userID.(string)

	todo, err := db.AddTodo(context, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task added successfully",
		"data":    todo,
	})
}

// UpdateTodoController, handler for todo put api
func UpdateTodoController(c *gin.Context) {
	context := c.Request.Context()
	todoID := c.Param("id")

	if strings.TrimSpace(todoID) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todo identifier required in url parameter for this request",
		})

		return
	}
	todo, err := db.GetTodoById(context, todoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	newTodo := &models.Todo{}
	if err := c.BindJSON(newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if newTodo.Title != "" && newTodo.Title != todo.Title {
		todo.Title = newTodo.Title
	}

	if newTodo.IsCompleted != todo.IsCompleted {
		todo.IsCompleted = newTodo.IsCompleted
	}

	todo, err = db.UpdateTodo(context, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"data":    todo,
	})
}

// DeleteTodoController, handler for todo delete api
func DeleteTodoController(c *gin.Context) {
	context := c.Request.Context()
	todoID := c.Param("id")

	if strings.TrimSpace(todoID) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todo identifier required in url parameter for this request",
		})

		return
	}
	_, err := db.GetTodoById(context, todoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	err = db.DeleteTodo(context, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}

// GetTodoController, handler to get todo as per userID
func GetTodoController(c *gin.Context) {
	context := c.Request.Context()
	userID, _ := c.Get("user_id")

	todos, err := db.GetTodos(context, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks listed successfully",
		"data":    todos,
	})
}
