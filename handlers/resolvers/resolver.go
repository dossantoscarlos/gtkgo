package resolvers

import (
	"gtkgo/handlers/actions"

	"github.com/gin-gonic/gin"
)

func ConfigRouters(g *gin.Engine) *gin.Engine {
	routers := g.Group("/api/v1")
	{
		users := routers.Group("users")
		{
			users.GET("/", actions.UserActionGetAll)
		}

		auth := routers.Group("/")
		{
			auth.POST("/register", actions.UserActionCreate)
			auth.POST("/login", actions.AuthenticationAction)
		}
	}

	return g
}
