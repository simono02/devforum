package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	// App
	AppEnv    string
	SecretKey string
	Debug     bool

	// Database
	DatabaseURL string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string

	// Mail (Brevo SMTP)
	MailServer        string
	MailPort          int
	MailUsername      string
	MailPassword      string
	MailUseTLS        bool
	MailUseSSL        bool
	MailDefaultSender string

	// App URL
	AppBaseURL string

	// Admin seed credentials
	AdminEmail    string
	AdminPassword string
	AdminName     string
}

// LoadConfig reads .env then environment variables and returns a populated Config
func LoadConfig() (*Config, error) {
	// Load .env file if present (ignore error — env vars may already be set)
	_ = godotenv.Load()

	cfg := &Config{
		// App
		AppEnv:    getEnv("APP_ENV", "development"),
		SecretKey: getEnv("SECRET_KEY", "dev-secret-key-change-in-production"),
		Debug:     getEnvBool("DEBUG", true),

		// Database — prefer full URL, fall back to individual vars
		DatabaseURL: getEnv("DATABASE_URL", ""),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "root"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBName:      getEnv("DB_NAME", "devforum_db"),

		// Mail
		MailServer:        getEnv("MAIL_SERVER", "smtp-relay.brevo.com"),
		MailPort:          getEnvInt("MAIL_PORT", 587),
		MailUsername:      getEnv("MAIL_USERNAME", ""),
		MailPassword:      getEnv("MAIL_PASSWORD", ""),
		MailUseTLS:        getEnvBool("MAIL_USE_TLS", true),
		MailUseSSL:        getEnvBool("MAIL_USE_SSL", false),
		MailDefaultSender: getEnv("MAIL_DEFAULT_SENDER", ""),

		// Base URL
		AppBaseURL: getEnv("APP_BASE_URL", "http://localhost:8080"),

		// Admin seed
		AdminEmail:    getEnv("ADMIN_EMAIL", "admin@devforum.com"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin123"),
		AdminName:     getEnv("ADMIN_NAME", "Admin User"),
	}

	// Build DSN
	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		)
	}

	return cfg, nil
}

// ── helpers ──────────────────────────────────────────────────────────────────

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b
		}
	}
	return fallback
}