package main

import (
	"fmt"
	"log"
	"net/http"

	"devforum/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	InitDB(cfg)
	SeedAdmin(cfg)

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS — allow Vue dev server
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// ── Routes ────────────────────────────────────────────
	routes.RegisterAuthRoutes(r, DB)
	routes.RegisterPostRoutes(r, DB)
	routes.RegisterUserRoutes(r, DB)
	routes.RegisterCommentRoutes(r, DB)

	addr := fmt.Sprintf(":%s", getEnv("PORT", "8080"))
	log.Printf("Server starting on %s (debug=%v)", addr, cfg.Debug)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}