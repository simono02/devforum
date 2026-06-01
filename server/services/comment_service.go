package services

import (
	"errors"

	"devforum/models"

	"gorm.io/gorm"
)

// ── Types ─────────────────────────────────────────────────────────────────────

type CreateCommentInput struct {
	Content string `json:"content" binding:"required,min=1"`
}

type UpdateCommentInput struct {
	Content string `json:"content" binding:"required,min=1"`
}

type CommentResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	PostID    uint   `json:"post_id"`
	Author    string `json:"author"`
	IsEdited  bool   `json:"is_edited"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ── Service ───────────────────────────────────────────────────────────────────

type CommentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{DB: db}
}

// GetByPost returns all comments for a post, oldest first.
func (s *CommentService) GetByPost(postID uint) ([]CommentResponse, error) {
	var comments []models.Comment
	if err := s.DB.Where("post_id = ?", postID).Order("created_at asc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return s.enrich(comments), nil
}

// Create adds a new comment to a post.
func (s *CommentService) Create(userID uint, postID uint, input CreateCommentInput) (*CommentResponse, error) {
	// Verify post exists
	var post models.Post
	if err := s.DB.First(&post, postID).Error; err != nil {
		return nil, errors.New("post not found")
	}

	comment := models.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  postID,
	}
	if err := s.DB.Create(&comment).Error; err != nil {
		return nil, errors.New("failed to create comment")
	}
	r := s.enrichOne(comment)
	return &r, nil
}

// Update edits a comment — owner only.
func (s *CommentService) Update(commentID uint, userID uint, input UpdateCommentInput) (*CommentResponse, error) {
	var comment models.Comment
	if err := s.DB.First(&comment, commentID).Error; err != nil {
		return nil, errors.New("comment not found")
	}
	if comment.UserID != userID {
		return nil, errors.New("you can only edit your own comments")
	}
	comment.Content  = input.Content
	comment.IsEdited = true
	if err := s.DB.Save(&comment).Error; err != nil {
		return nil, errors.New("failed to update comment")
	}
	r := s.enrichOne(comment)
	return &r, nil
}

// Delete removes a comment — owner or admin only.
func (s *CommentService) Delete(commentID uint, userID uint, role string) error {
	var comment models.Comment
	if err := s.DB.First(&comment, commentID).Error; err != nil {
		return errors.New("comment not found")
	}
	if comment.UserID != userID && role != "admin" {
		return errors.New("you can only delete your own comments")
	}
	return s.DB.Delete(&comment).Error
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func (s *CommentService) enrichOne(c models.Comment) CommentResponse {
	var user models.User
	s.DB.Select("full_name").First(&user, c.UserID)
	return CommentResponse{
		ID:        c.ID,
		Content:   c.Content,
		UserID:    c.UserID,
		PostID:    c.PostID,
		Author:    user.FullName,
		IsEdited:  c.IsEdited,
		CreatedAt: c.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: c.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func (s *CommentService) enrich(comments []models.Comment) []CommentResponse {
	out := make([]CommentResponse, len(comments))
	for i, c := range comments {
		out[i] = s.enrichOne(c)
	}
	return out
}