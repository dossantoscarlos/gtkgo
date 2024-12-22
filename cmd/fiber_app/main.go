package main

import (
	"gtkgo/fiber_actions/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.ConfigRouters(app)
	app.Listen(":3000")

}
