package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/tweet"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user/follow"
)

func TestCreateRepositories(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expect := &repositories.Repositories{
		UserRepository:   user.NewRepository(db),
		TweetRepository:  tweet.NewRepository(db),
		FollowRepository: follow.NewRepository(db),
	}

	actual := repositories.CreateRepositories(db)

	assert.Equal(t, expect, actual)
}
