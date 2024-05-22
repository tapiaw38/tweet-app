package user

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user/follow"
	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	FollowedUsecase interface {
		Execute(context.Context, domain.Follow) error
	}

	followedUsecase struct {
		repository follow.Repository
	}
)

func NewFollowedUsecase(repository follow.Repository) FollowedUsecase {
	return &followedUsecase{
		repository: repository,
	}
}

func (u *followedUsecase) Execute(ctx context.Context, follow domain.Follow) error {
	err := u.repository.Create(ctx, follow)

	return err
}
