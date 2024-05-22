package user

import (
	"context"
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)


func (r *repository) Create(ctx context.Context, user domain.User) (uint64, error) {
	insertedID, err := r.executeCreateQuery(ctx, user)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *repository) executeCreateQuery(ctx context.Context, user domain.User) (uint64, error) {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	args := []any{
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
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