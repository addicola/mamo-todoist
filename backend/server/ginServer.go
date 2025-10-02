package server

import "github.com/gin-gonic/gin"

type GinServer struct {
	app *gin.Engine
}

func NewGinServer() Server {
	app := gin.Default()
	return GinServer{
		app: app,
	}
}

func (s GinServer) App() *gin.Engine {
	return s.app
}

func (s GinServer) Start() {
	s.App().GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.App().Run(":8080")
}
