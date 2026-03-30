package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DatabaseURL   string
	SupabaseDBURL string
	AllowedOrigin string
}

func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		Port:          getEnv("PORT", "8080"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		SupabaseDBURL: os.Getenv("SUPABASE_DB_URL"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "*"),
	}

	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = cfg.SupabaseDBURL
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL or SUPABASE_DB_URL is required")
	}

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
