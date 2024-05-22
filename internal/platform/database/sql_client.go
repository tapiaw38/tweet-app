package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/tapiaw38/tweet-app/internal/platform/config"
)

var sqlClient *sql.DB

type SQLConfig struct {
	DatabaseURL string
}

func NewSQLConfig(config config.Config) SQLConfig {
	return SQLConfig{
		DatabaseURL: config.DatabaseURL,
	}
}

func (c *SQLConfig) initSQLClient() error {
	newSQLClient, err := sql.Open("postgres", c.DatabaseURL)
	if err != nil {
		return err
	}

	if err = newSQLClient.Ping(); err != nil {
		return err
	}

	sqlClient = newSQLClient
	return nil
}

func (c *SQLConfig) GetSQLClientInstance() (*sql.DB, error) {
	if sqlClient == nil {
		if err := c.initSQLClient(); err != nil {
			return nil, err
		}
	}
	return sqlClient, nil
}

func getRelativePathToMigrationsDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	absMigrationsDirPath := filepath.Join(cwd, "migrations")

	relMigrationsDirPath, err := filepath.Rel(cwd, absMigrationsDirPath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("file://%s", relMigrationsDirPath), nil
}

func (c *SQLConfig) Makemigration() error {
	migrationPath, err := getRelativePathToMigrationsDirectory()
	if err != nil {
		return err
	}

	m, err := migrate.New(migrationPath, c.DatabaseURL)
	if err != nil {
		return err
	}

	version, _, _ := m.Version()
	log.Printf("migrations: current version is %v", version)

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("migrations:", err)
			return nil
		}
		return err
	}

	log.Println("migrations: database migrated")

	return nil
}
