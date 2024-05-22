package follow

import (
	"context"
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

func (r *repository) Create(ctx context.Context, follow domain.Follow) error {
	err := r.executeCreateQuery(ctx, follow)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) executeCreateQuery(ctx context.Context, follow domain.Follow) error {
	query := `
		INSERT INTO follows (user_id, follow_id, created_at)
		VALUES ($1, $2, $3)
	`

	args := []any{
		follow.UserID,
		follow.FollowID,
		time.Now(),
	}

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
