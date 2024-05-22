package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/tweet-app/internal/usecases/user"
)

func NewCreateHandler(usecase user.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest UserInputData
		if err := c.BindJSON(&userRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertUser := toUserInput(userRequest)
		createdUser, err := usecase.Execute(c.Request.Context(), insertUser)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, createdUser)
	}
}