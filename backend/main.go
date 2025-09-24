package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("todos", listTodos)
	r.GET("todos/:id", getTodo)

	r.Run(":8080")
}

func listTodos(c *gin.Context) {
	// Placeholder for listing todos
	c.JSON(200, gin.H{
		"message": "List of todos",
	})
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	// Placeholder for getting a specific todo by id
	c.JSON(200, gin.H{
		"message": "Get todo with ID " + id,
	})
}
