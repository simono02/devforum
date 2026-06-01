package main

import (
	"log"

	"devforum/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database instance (analogous to Flask-SQLAlchemy's `db`)
var DB *gorm.DB

// InitDB opens the Postgres connection and runs auto-migrations.
// Call this once from main() after loading config.
func InitDB(cfg *Config) {
	var err error

	gormCfg := &gorm.Config{}
	if cfg.Debug {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Silent)
	}

	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), gormCfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")

	// Run migrations — equivalent to `flask db upgrade`
	// Add every model struct here as you create them.
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	log.Println("Database migrations completed")
}

// SeedAdmin creates the admin user from .env if they don't exist yet.
func SeedAdmin(cfg *Config) {
	if cfg.AdminEmail == "" || cfg.AdminPassword == "" {
		return
	}

	var existing models.User
	if err := DB.Where("email = ?", cfg.AdminEmail).First(&existing).Error; err == nil {
		// Already exists — make sure role is admin
		if existing.Role != "admin" {
			DB.Model(&existing).Update("role", "admin")
			log.Printf("Updated %s to admin role", cfg.AdminEmail)
		} else {
			log.Printf("Admin user already exists: %s", cfg.AdminEmail)
		}
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash admin password: %v", err)
	}

	admin := models.User{
		FullName: cfg.AdminName,
		Username: "admin",
		Email:    cfg.AdminEmail,
		Password: string(hashed),
		Role:     "admin",
		IsActive: true,
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Printf("Failed to seed admin user: %v", err)
		return
	}

	log.Printf("Admin user created: %s", cfg.AdminEmail)
}