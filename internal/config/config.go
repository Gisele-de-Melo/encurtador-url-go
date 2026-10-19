package config

import (
	"os"
)

type Config struct {
	Port         string
	Env          string
	DatabasePath string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:         getEnv("APP_PORT", "8080"),
		Env:          getEnv("APP_ENV", "development"),
		DatabasePath: getEnv("DATABASE_PATH", "shortener.db"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
