package repositories

import (
	"database/sql"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/tweet"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user/follow"
)

type Repositories struct {
	UserRepository   user.Repository
	TweetRepository  tweet.Repository
	FollowRepository follow.Repository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository:   user.NewRepository(db),
		TweetRepository:  tweet.NewRepository(db),
		FollowRepository: follow.NewRepository(db),
	}
}
