package tweet

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/tweet"
	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.Tweet) (CreateOutput, error)
	}

	createUsecase struct {
		repository tweet.Repository
	}

	CreateOutput struct {
		Data TweetOutputData `json:"data"`
	}
)

func NewCreateUsecase(repository tweet.Repository) CreateUsecase {
	return &createUsecase{
		repository: repository,
	}
}

func (u *createUsecase) Execute(ctx context.Context, tweet domain.Tweet) (CreateOutput, error) {
	tweetID, err := u.repository.Create(ctx, tweet)
	if err != nil {
		return CreateOutput{}, err
	}

	tweetCreated, err := u.repository.Get(ctx, tweetID)
	if err != nil || tweetCreated == nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		Data: toTweetOutput(*tweetCreated),
	}, nil
}
