package routes

import (
	"net/http"
	"strconv"

	"devforum/middleware"
	"devforum/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPostRoutes(r *gin.Engine, db *gorm.DB) {
	svc := services.NewPostService(db)

	posts := r.Group("/api/posts")

	// ── Public ────────────────────────────────────────────

	// GET /api/posts
	posts.GET("", func(c *gin.Context) {
		list, err := svc.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	})

	// GET /api/posts/:id  (must come AFTER /mine)
	posts.GET("/:id", func(c *gin.Context) {
		if c.Param("id") == "mine" {
			// delegate to mine handler with auth check
			c.Next()
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
			return
		}
		post, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	// ── Authenticated ─────────────────────────────────────

	// GET /api/posts/mine
	posts.GET("/mine", middleware.AuthRequired(), func(c *gin.Context) {
		userID := c.GetUint("user_id")
		list, err := svc.GetByUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	})

	// POST /api/posts — create
	posts.POST("", middleware.AuthRequired(), func(c *gin.Context) {
		var input services.CreatePostInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.GetUint("user_id")
		post, err := svc.Create(userID, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, post)
	})

	// PUT /api/posts/:id — update
	posts.PUT("/:id", middleware.AuthRequired(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
			return
		}
		var input services.UpdatePostInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.GetUint("user_id")
		post, err := svc.Update(uint(id), userID, input)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	// DELETE /api/posts/:id
	posts.DELETE("/:id", middleware.AuthRequired(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
			return
		}
		userID := c.GetUint("user_id")
		role, _ := c.Get("role")
		if err := svc.Delete(uint(id), userID, role.(string)); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
	})

	// POST /api/posts/:id/upvote
	posts.POST("/:id/upvote", middleware.AuthRequired(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
			return
		}
		post, err := svc.Upvote(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	})
}