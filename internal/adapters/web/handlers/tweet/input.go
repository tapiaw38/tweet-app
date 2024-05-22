package tweet

import (
	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	TweetInputData struct {
		UserID  uint64 `json:"user_id"`
		Content string `json:"content"`
	}
)

func toTweetInput(tweet TweetInputData) domain.Tweet {
	return domain.Tweet{
		UserID:  tweet.UserID,
		Content: tweet.Content,
	}
}
