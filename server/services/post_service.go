package services

import (
	"errors"
	"time"

	"devforum/models"

	"gorm.io/gorm"
)

// ── Input / Response types ────────────────────────────────────────────────────

type CreatePostInput struct {
	Title   string `json:"title"   binding:"required,min=5"`
	Content string `json:"content" binding:"required,min=10"`
	Tag     string `json:"tag"`
}

type UpdatePostInput struct {
	Title   string `json:"title"   binding:"required,min=5"`
	Content string `json:"content" binding:"required,min=10"`
	Tag     string `json:"tag"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tag       string    `json:"tag"`
	UserID    uint      `json:"user_id"`
	Author    string    `json:"author"`
	Upvotes   int       `json:"upvotes"`
	Comments  int64     `json:"comments"`
	IsPinned  bool      `json:"is_pinned"`
	IsFlagged bool      `json:"is_flagged"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ── Service ───────────────────────────────────────────────────────────────────

type PostService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{DB: db}
}

// GetAll returns all posts newest first, each enriched with author name.
func (s *PostService) GetAll() ([]PostResponse, error) {
	var posts []models.Post
	if err := s.DB.Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return s.enrichPosts(posts), nil
}

// GetByID returns a single post by ID.
func (s *PostService) GetByID(id uint) (*PostResponse, error) {
	var post models.Post
	if err := s.DB.First(&post, id).Error; err != nil {
		return nil, errors.New("post not found")
	}
	r := s.enrichPost(post)
	return &r, nil
}

// GetByUser returns all posts belonging to a specific user.
func (s *PostService) GetByUser(userID uint) ([]PostResponse, error) {
	var posts []models.Post
	if err := s.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return s.enrichPosts(posts), nil
}

// Create saves a new post and returns it.
func (s *PostService) Create(userID uint, input CreatePostInput) (*PostResponse, error) {
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		Tag:     input.Tag,
		UserID:  userID,
	}
	if err := s.DB.Create(&post).Error; err != nil {
		return nil, errors.New("failed to create post")
	}
	r := s.enrichPost(post)
	return &r, nil
}

// Update edits a post — only the owner may update.
func (s *PostService) Update(postID uint, userID uint, input UpdatePostInput) (*PostResponse, error) {
	var post models.Post
	if err := s.DB.First(&post, postID).Error; err != nil {
		return nil, errors.New("post not found")
	}
	if post.UserID != userID {
		return nil, errors.New("you can only edit your own posts")
	}
	post.Title   = input.Title
	post.Content = input.Content
	post.Tag     = input.Tag
	if err := s.DB.Save(&post).Error; err != nil {
		return nil, errors.New("failed to update post")
	}
	r := s.enrichPost(post)
	return &r, nil
}

// Delete removes a post — only the owner or an admin may delete.
func (s *PostService) Delete(postID uint, userID uint, role string) error {
	var post models.Post
	if err := s.DB.First(&post, postID).Error; err != nil {
		return errors.New("post not found")
	}
	if post.UserID != userID && role != "admin" {
		return errors.New("you can only delete your own posts")
	}
	return s.DB.Delete(&post).Error
}

// Upvote increments the upvote counter.
func (s *PostService) Upvote(postID uint) (*PostResponse, error) {
	var post models.Post
	if err := s.DB.First(&post, postID).Error; err != nil {
		return nil, errors.New("post not found")
	}
	post.Upvotes++
	if err := s.DB.Save(&post).Error; err != nil {
		return nil, errors.New("failed to upvote post")
	}
	r := s.enrichPost(post)
	return &r, nil
}

// ── Helpers ───────────────────────────────────────────────────────────────────

// enrichPost joins author full_name and comment count from the DB.
func (s *PostService) enrichPost(p models.Post) PostResponse {
	var user models.User
	s.DB.Select("full_name").First(&user, p.UserID)
	var commentCount int64
	s.DB.Model(&models.Comment{}).Where("post_id = ?", p.ID).Count(&commentCount)
	return PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		Tag:       p.Tag,
		UserID:    p.UserID,
		Author:    user.FullName,
		Upvotes:   p.Upvotes,
		Comments:  commentCount,
		IsPinned:  p.IsPinned,
		IsFlagged: p.IsFlagged,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (s *PostService) enrichPosts(posts []models.Post) []PostResponse {
	out := make([]PostResponse, len(posts))
	for i, p := range posts {
		out[i] = s.enrichPost(p)
	}
	return out
}