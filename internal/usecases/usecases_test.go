package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories"
	tweet_repository "github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/tweet/mocks"
	follow_repository "github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user/follow/mocks"
	user_repository "github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user/mocks"
	"github.com/tapiaw38/tweet-app/internal/usecases"
	"github.com/tapiaw38/tweet-app/internal/usecases/tweet"
	"github.com/tapiaw38/tweet-app/internal/usecases/user"
	"go.uber.org/mock/gomock"
)

func TestCreateRepositories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositories := &repositories.Repositories{
		UserRepository:   user_repository.NewMockRepository(ctrl),
		TweetRepository:  tweet_repository.NewMockRepository(ctrl),
		FollowRepository: follow_repository.NewMockRepository(ctrl),
	}

	expect := &usecases.UseCases{
		CreateUserUseCase:  user.NewCreateUsecase(repositories.UserRepository),
		GetUserUseCase:     user.NewGetUsecase(repositories.UserRepository),
		FollowedUseCase:    user.NewFollowedUsecase(repositories.FollowRepository),
		CreateTweetUseCase: tweet.NewCreateUsecase(repositories.TweetRepository),
		GetTweetUseCase:    tweet.NewGetUsecase(repositories.TweetRepository),
		ListTweetUseCase:   tweet.NewListUsecase(repositories.TweetRepository),
	}

	actual := usecases.CreateUsecases(repositories)

	assert.Equal(t, expect, actual)
}
