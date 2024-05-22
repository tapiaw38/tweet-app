package tweet

import (
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	TweetOutputData struct {
		ID        uint64          `json:"id"`
		Content   string          `json:"content"`
		User      *UserOutputData `json:"user"`
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
	}

	UserOutputData struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Username  string `json:"username,omitempty"`
		Email     string `json:"email,omitempty"`
	}
)

func toTweetOutput(tweet domain.Tweet) TweetOutputData {
	return TweetOutputData{
		ID:      tweet.ID,
		Content: tweet.Content,
		User: &UserOutputData{
			FirstName: tweet.User.FirstName,
			LastName:  tweet.User.LastName,
			Username:  tweet.User.Username,
			Email:     tweet.User.Email,
		},
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
}

func toTweetOutputList(tweets []domain.Tweet) []TweetOutputData {
	tweetOutputList := []TweetOutputData{}

	for _, tweet := range tweets {
		tweetOutputList = append(tweetOutputList, toTweetOutput(tweet))
	}

	return tweetOutputList
}
