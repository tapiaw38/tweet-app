package user

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	Repository interface {
		Create(context.Context, domain.User) (uint64, error)
		Get(context.Context, uint64) (*domain.User, error)
		GetByUsername(context.Context, string) (*domain.User, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
