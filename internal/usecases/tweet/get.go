package tweet

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/tweet"
)

type (
	GetUsecase interface {
		Execute(context.Context, uint64) (*GetOutput, error)
	}

	getUsecase struct {
		repository tweet.Repository
	}

	GetOutput struct {
		Data TweetOutputData `json:"data"`
	}
)

func NewGetUsecase(repository tweet.Repository) GetUsecase {
	return &getUsecase{
		repository: repository,
	}
}

func (u *getUsecase) Execute(ctx context.Context, id uint64) (*GetOutput, error) {
	tweet, err := u.repository.Get(ctx, id)
	if err != nil || tweet == nil {
		return nil, err
	}

	return &GetOutput{
		Data: toTweetOutput(*tweet),
	}, nil
}
