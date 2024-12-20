package main

import (
	"gtkgo/handlers/resolvers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	server := resolvers.ConfigRouters(router)

	server.Run(":8080")
}
