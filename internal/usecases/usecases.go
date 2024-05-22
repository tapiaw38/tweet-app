package usecases

import (
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/tweet-app/internal/usecases/tweet"
	"github.com/tapiaw38/tweet-app/internal/usecases/user"
)

type UseCases struct {
	CreateUserUseCase  user.CreateUsecase
	GetUserUseCase     user.GetUsecase
	CreateTweetUseCase tweet.CreateUsecase
	GetTweetUseCase    tweet.GetUsecase
	ListTweetUseCase   tweet.ListUsecase
	FollowedUseCase    user.FollowedUsecase
}

func CreateUsecases(repositories *repositories.Repositories) *UseCases {
	return &UseCases{
		CreateUserUseCase:  user.NewCreateUsecase(repositories.UserRepository),
		GetUserUseCase:     user.NewGetUsecase(repositories.UserRepository),
		FollowedUseCase:    user.NewFollowedUsecase(repositories.FollowRepository),
		CreateTweetUseCase: tweet.NewCreateUsecase(repositories.TweetRepository),
		GetTweetUseCase:    tweet.NewGetUsecase(repositories.TweetRepository),
		ListTweetUseCase:   tweet.NewListUsecase(repositories.TweetRepository),
	}
}
