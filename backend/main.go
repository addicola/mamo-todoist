package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("todos", listTodos)

	r.Run(":8080")
}

func listTodos(c *gin.Context) {
	// Placeholder for listing todos
	c.JSON(200, gin.H{
		"message": "List of todos",
	})
}
