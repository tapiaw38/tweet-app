package config

import "os"

type Config struct {
	GinMode     string
	Port        string
	Host        string
	DatabaseURL string
}

func NewConfig() Config {
	return Config{
		GinMode:     getEnv("GIN_MODE", "debug"),
		Port:        getEnv("PORT", "8080"),
		Host:        getEnv("HOST", "localhost"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
