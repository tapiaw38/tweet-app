package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/tweet-app/internal/platform/config"
	"github.com/tapiaw38/tweet-app/internal/platform/database"
)

func TestNewSQLConfig(t *testing.T) {
	cfg := config.Config{DatabaseURL: "postgres://localhost/testdb"}
	sqlConfig := database.NewSQLConfig(cfg)

	assert.Equal(t, cfg.DatabaseURL, sqlConfig.DatabaseURL)
}
