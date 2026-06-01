package routes

import (
	"net/http"

	"devforum/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAuthRoutes mounts /api/auth/register and /api/auth/login.
func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	svc := services.NewAuthService(db)
	auth := r.Group("/api/auth")

	// POST /api/auth/register
	auth.POST("/register", func(c *gin.Context) {
		var input services.RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := svc.Register(input)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	})

	// POST /api/auth/login
	auth.POST("/login", func(c *gin.Context) {
		var input services.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := svc.Login(input)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}