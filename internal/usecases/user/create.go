package user

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.User) (CreateOutput, error)
	}

	createUsecase struct {
		repository user.Repository
	}

	CreateOutput struct {
		Data UserOutputData `json:"data"`
	}
)

func NewCreateUsecase(repository user.Repository) CreateUsecase {
	return &createUsecase{
		repository: repository,
	}
}

func (u *createUsecase) Execute(ctx context.Context, user domain.User) (CreateOutput, error) {
	userID, err := u.repository.Create(ctx, user)
	if err != nil {
		return CreateOutput{}, err
	}

	userCreated, err := u.repository.Get(ctx, userID)
	if err != nil || userCreated == nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		Data: toUserOutput(*userCreated),
	}, nil
}