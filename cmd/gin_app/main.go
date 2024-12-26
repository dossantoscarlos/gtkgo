package main

import (
	"gtkgo/client/gin_actions/resolvers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	server := resolvers.ConfigRouters(router)

	server.Run(":8080")
}
