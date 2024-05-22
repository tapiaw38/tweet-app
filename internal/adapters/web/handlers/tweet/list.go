package tweet

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/tweet-app/internal/usecases/tweet"
)

func NewListHandler(usecase tweet.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := parseListFilterOptions(c.Request.URL.Query())

		tweets, err := usecase.Execute(c.Request.Context(), filters)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, tweets)
	}
}

func parseListFilterOptions(queries url.Values) tweet.ListFilterOptions {
	return tweet.ListFilterOptions{
		Username: queries.Get("username"),
	}
}
