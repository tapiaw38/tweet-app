package tweet

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/tweet-app/internal/usecases/tweet"
)

func NewCreateHandler(usecase tweet.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tweetRequest TweetInputData
		if err := c.BindJSON(&tweetRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertTweet := toTweetInput(tweetRequest)
		tweet, err := usecase.Execute(c.Request.Context(), insertTweet)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, tweet)
	}
}
