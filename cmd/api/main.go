package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/tweet-app/internal/adapters/web"
	"github.com/tapiaw38/tweet-app/internal/platform/config"
	"github.com/tapiaw38/tweet-app/internal/platform/database"
	"github.com/tapiaw38/tweet-app/internal/usecases"
)

func main() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}

	if err := run(config); err != nil {
		panic(err)
	}
}

func initConfig() (*config.Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error load env file")
	}

	cfg := config.NewConfig()
	if cfg.DatabaseURL == "" {
		return nil, errors.New("databaseURL is required")
	}
	if cfg.Port == "" {
		return nil, errors.New("port is required")
	}

	return &cfg, nil
}

func run(config *config.Config) error {
	sqlClient := database.NewSQLConfig(*config)

	db, err := sqlClient.GetSQLClientInstance()
	if err != nil {
		return err
	}
	defer db.Close()

	err = sqlClient.Makemigration()
	if err != nil {
		return err
	}

	if config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()

	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"*"}
	cfg.AllowHeaders = []string{"*"}
	cfg.ExposeHeaders = []string{"*"}

	app.Use(cors.New(cfg))

	bootstrap(app, db)

	return app.Run(":" + config.Port)
}

func bootstrap(app *gin.Engine, db *sql.DB) {
	repositories := repositories.CreateRepositories(db)
	usecases := usecases.CreateUsecases(repositories)

	web.RegisterApplicationRoutes(app, usecases)
}
