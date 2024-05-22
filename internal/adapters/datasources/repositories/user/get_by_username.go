package user

import (
	"context"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

func (r *repository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := r.executeGetByUsernameQuery(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) executeGetByUsernameQuery(ctx context.Context, username string) (*domain.User, error) {
	query := `
		SELECT id, first_name, last_name, email, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	args := []any{
		username,
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
