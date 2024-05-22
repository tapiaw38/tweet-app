package user

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user"
)


type (
	GetUsecase interface {
		Execute(context.Context, uint64) (*GetOutput, error)
	}

	getUsecase struct {
		repository user.Repository
	}

	GetOutput struct {
		Data UserOutputData `json:"data"`
	}
)

func NewGetUsecase(repository user.Repository) GetUsecase {
	return &getUsecase{
		repository: repository,
	}
}

func (u *getUsecase) Execute(ctx context.Context, id uint64) (*GetOutput, error) {
	user, err := u.repository.Get(ctx, id)
	if err != nil || user == nil {
		return nil, err
	}

	return &GetOutput{
		Data: toUserOutput(*user),
	}, nil
}