package tweet

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	Repository interface {
		Create(context.Context, domain.Tweet) (uint64, error)
		Get(context.Context, uint64) (*domain.Tweet, error)
		List(context.Context, ListFilterOptions) ([]domain.Tweet, error)
	}

	repository struct {
		db *sql.DB
	}

	ListFilterOptions struct {
		Username string
	}
)

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
