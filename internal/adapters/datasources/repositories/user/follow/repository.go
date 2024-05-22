package follow

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

type (
	Repository interface {
		Create(ctx context.Context, follow domain.Follow) error
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
