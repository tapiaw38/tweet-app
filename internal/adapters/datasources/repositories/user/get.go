package user

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/domain"
)


func (r *repository) Get(ctx context.Context, id uint64) (*domain.User, error) {
	user, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) executeGetQuery(ctx context.Context, id uint64) (*domain.User, error) {
	query := `
		SELECT id, first_name, last_name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	args := []any{
		id,
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	user := domain.User{}

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}