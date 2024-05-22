package tweet

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/tweet"
)

type (
	ListUsecase interface {
		Execute(context.Context, ListFilterOptions) (ListOutput, error)
	}

	listUsecase struct {
		repository tweet.Repository
	}

	ListFilterOptions tweet.ListFilterOptions

	ListOutput struct {
		Data []TweetOutputData `json:"data"`
	}
)

func NewListUsecase(repository tweet.Repository) ListUsecase {
	return &listUsecase{
		repository: repository,
	}
}

func (u *listUsecase) Execute(ctx context.Context, filter ListFilterOptions) (ListOutput, error) {
	var tweetsData ListOutput

	teets, err := u.repository.List(ctx, tweet.ListFilterOptions(filter))
	if err != nil {
		return tweetsData, err
	}

	tweetsData = ListOutput{
		Data: toTweetOutputList(teets),
	}

	return tweetsData, nil
}
