package tweet

import (
	"context"
	"time"

	"github.com/tapiaw38/tweet-app/internal/domain"
)

func (r *repository) List(ctx context.Context, filters ListFilterOptions) ([]domain.Tweet, error) {
	tweets, err := r.executeListQuery(ctx, filters)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}

func (r *repository) executeListQuery(ctx context.Context, filters ListFilterOptions) ([]domain.Tweet, error) {
	query := `SELECT tw.id, us.first_name, us.last_name, us.username, us.email,
				tw.content, tw.created_at, tw.updated_at
			FROM tweets tw
			LEFT JOIN users us ON tw.user_id = us.id
	`

	query += ` WHERE tw.id = tw.id`

	var args []any

	if filters.Username != "" {
		query += ` AND tw.user_id IN (
			SELECT fw.follow_id FROM follows fw WHERE fw.user_id = (
				SELECT id FROM users WHERE username = $1
			)
		)`
		args = append(args, filters.Username)
	}

	query += ` ORDER BY tw.created_at DESC`
	query += ` ;`

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tweets := []domain.Tweet{}
	for rows.Next() {
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
		err = rows.Scan(
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

		tweets = append(tweets, unmarshalTweet(
			tweetID,
			firstName,
			lastName,
			username,
			email,
			content,
			createdAt,
			updatedAt,
		))
	}

	return tweets, nil
}
