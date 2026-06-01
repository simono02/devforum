package routes

import (
	"net/http"
	"strconv"

	"devforum/middleware"
	"devforum/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCommentRoutes(r *gin.Engine, db *gorm.DB) {
	svc := services.NewCommentService(db)

	// GET /api/posts/:id/comments — public
	r.GET("/api/posts/:id/comments", func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
			return
		}
		list, err := svc.GetByPost(uint(postID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	})

	// POST /api/posts/:id/comments — auth required
	r.POST("/api/posts/:id/comments", middleware.AuthRequired(), func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
			return
		}
		var input services.CreateCommentInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.GetUint("user_id")
		comment, err := svc.Create(userID, uint(postID), input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, comment)
	})

	// PUT /api/comments/:id — auth required, owner only
	r.PUT("/api/comments/:id", middleware.AuthRequired(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
			return
		}
		var input services.UpdateCommentInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.GetUint("user_id")
		comment, err := svc.Update(uint(id), userID, input)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comment)
	})

	// DELETE /api/comments/:id — auth required, owner or admin
	r.DELETE("/api/comments/:id", middleware.AuthRequired(), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
			return
		}
		userID := c.GetUint("user_id")
		role, _ := c.Get("role")
		if err := svc.Delete(uint(id), userID, role.(string)); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
	})
}