package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/tweet-app/internal/adapters/web/handlers/tweet"
	"github.com/tapiaw38/tweet-app/internal/adapters/web/handlers/user"
	"github.com/tapiaw38/tweet-app/internal/usecases"
)

func RegisterApplicationRoutes(app *gin.Engine, usecases *usecases.UseCases) {

	routeGroup := app.Group("/api")

	routeGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routeGroup.POST("/users", user.NewCreateHandler(usecases.CreateUserUseCase))
	routeGroup.POST("/users/follow/:id", user.NewFollowerHandler(usecases.FollowedUseCase))

	routeGroup.POST("/tweets", tweet.NewCreateHandler(usecases.CreateTweetUseCase))
	routeGroup.GET("/tweets", tweet.NewListHandler(usecases.ListTweetUseCase))
}
