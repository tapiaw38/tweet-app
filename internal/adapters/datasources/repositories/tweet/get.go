package tweet

import (
	"context"
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

func (r repository) Get(ctx context.Context, id uint64) (*domain.Tweet, error) {
	tweet, err := r.executeGetQuery(ctx, id)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (r repository) executeGetQuery(ctx context.Context, id uint64) (*domain.Tweet, error) {
	query := `
		SELECT tw.id,
			us.first_name, us.last_name, us.username, us.email,
			tw.content, tw.created_at, tw.updated_at
		FROM tweets tw
		LEFT JOIN users us ON tw.user_id = us.id
		WHERE tw.id = $1;
	`

	args := []any{
		id,
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var (
		tweetID   uint64
		firstName string
		lastName  string
		username  string
		email     string
		content   string
		createdAt time.Time
		updatedAt time.Time
	)

	tweet := domain.Tweet{}

	err := row.Scan(
		&tweetID,
		&firstName,
		&lastName,
		&username,
		&email,
		&content,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	tweet = unmarshalTweet(
		tweetID,
		firstName,
		lastName,
		username,
		email,
		content,
		createdAt,
		updatedAt,
	)

	return &tweet, nil
}
