package server

import (
	"backend/config"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	conf *config.Config
	app  *gin.Engine
}

func NewGinServer(conf *config.Config) Server {
	app := gin.Default()
	return GinServer{
		conf: conf,
		app:  app,
	}
}

func (s GinServer) App() *gin.Engine {
	return s.app
}

func (s GinServer) Port() string {
	return s.conf.Server.Port
}

func (s GinServer) Start() {
	s.App().GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.App().Run(s.Port())
}
