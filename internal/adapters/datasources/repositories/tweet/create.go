package tweet

import (
	"context"
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

func (r *repository) Create(ctx context.Context, tweet domain.Tweet) (uint64, error) {
	insertedID, err := r.executeCreateQuery(ctx, tweet)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *repository) executeCreateQuery(ctx context.Context, tweet domain.Tweet) (uint64, error) {
	query := `
		INSERT INTO tweets (user_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	args := []any{
		tweet.UserID,
		tweet.Content,
		time.Now(),
		time.Now(),
	}

	row := r.db.QueryRowContext(
		ctx,
		query,
		args...,
	)

	var lastID uint64

	err := row.Scan(&lastID)

	return lastID, err
}
