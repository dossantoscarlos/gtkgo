package routes

import (
	"gtkgo/client/fiber_actions/actions"

	"github.com/gofiber/fiber/v2"
)

func ConfigRouters(f *fiber.App) *fiber.App {
	routers := f.Group("/api/v1")
	{
		users := routers.Group("users")
		{
			users.Get("/", actions.UserActionGetAll)
			users.Get("/show", actions.GetOneUsers)
			users.Delete("/delete", actions.UserActionDelete)
		}

		auth := routers.Group("/")
		{
			auth.Post("/register", actions.UserActionCreate)
			auth.Post("/login", actions.AuthenticationAction)
		}

	}

	return f
}
