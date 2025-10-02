package main

import (
	"backend/server"
)

func main() {
	server.NewGinServer().Start()
}
