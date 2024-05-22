package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/tweet-app/internal/usecases/user"
)

func NewFollowerHandler(usecase user.FollowedUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDParam := c.Params.ByName("id")
		if userIDParam == "" {
			c.JSON(400, gin.H{"error": "user id is required"})
			return
		}

		var followRequest FollowInputData
		if err := c.BindJSON(&followRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userID, _ := strconv.ParseUint(userIDParam, 10, 64)
		followRequest.UserID = userID

		insertFollow := toFollowInput(followRequest)
		err := usecase.Execute(c.Request.Context(), insertFollow)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "followed add successfully"})
	}
}
