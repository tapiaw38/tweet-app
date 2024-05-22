package tweet

import (
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

func unmarshalTweet(
	tweetID uint64,
	firstName string,
	lastName string,
	username string,
	email string,
	content string,
	createdAt time.Time,
	updatedAt time.Time,
) domain.Tweet {
	return domain.Tweet{
		ID: tweetID,
		User: &domain.User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Email:     email,
		},
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
