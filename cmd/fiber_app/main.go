package main

import (
	"gtkgo/client/fiber_actions/routes"

	_ "gtkgo/cmd/fiber_app/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	app := fiber.New()

	routes.ConfigRouters(app)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
