package resolvers

import (
	"gtkgo/client/gin_actions/actions"

	"github.com/gin-gonic/gin"
)

func ConfigRouters(g *gin.Engine) *gin.Engine {
	routers := g.Group("/api/v1")
	{
		users := routers.Group("users")
		{
			users.GET("/", actions.UserActionGetAll)
			users.GET("/show", actions.GetOneUsers)
		}

		auth := routers.Group("/")
		{
			auth.POST("/register", actions.UserActionCreate)
			auth.POST("/login", actions.AuthenticationAction)
		}
	}

	return g
}
