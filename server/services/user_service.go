package services

import (
	"errors"

	"devforum/models"

	"gorm.io/gorm"
)

// ── Types ─────────────────────────────────────────────────────────────────────

type UserResponse struct {
	ID        uint   `json:"id"`
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	Posts     int64  `json:"posts"`
}

type UpdateUserInput struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

// ── Service ───────────────────────────────────────────────────────────────────

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// GetAll returns all users with their post counts.
func (s *UserService) GetAll() ([]UserResponse, error) {
	var users []models.User
	if err := s.DB.Order("created_at desc").Find(&users).Error; err != nil {
		return nil, err
	}

	out := make([]UserResponse, len(users))
	for i, u := range users {
		var count int64
		s.DB.Model(&models.Post{}).Where("user_id = ?", u.ID).Count(&count)
		out[i] = UserResponse{
			ID:        u.ID,
			FullName:  u.FullName,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z"),
			Posts:     count,
		}
	}
	return out, nil
}

// GetStats returns platform-wide counts for the admin dashboard.
func (s *UserService) GetStats() (map[string]int64, error) {
	stats := map[string]int64{}

	var userCount, postCount int64
	s.DB.Model(&models.User{}).Count(&userCount)
	s.DB.Model(&models.Post{}).Count(&postCount)

	stats["users"] = userCount
	stats["posts"] = postCount
	return stats, nil
}

// Update saves changes to a user record.
func (s *UserService) Update(id uint, input UpdateUserInput) (*UserResponse, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	user.FullName = input.FullName
	user.Username = input.Username
	user.Email    = input.Email
	user.Role     = input.Role
	user.IsActive = input.IsActive
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}
	var count int64
	s.DB.Model(&models.Post{}).Where("user_id = ?", user.ID).Count(&count)
	resp := UserResponse{
		ID: user.ID, FullName: user.FullName, Username: user.Username,
		Email: user.Email, Role: user.Role, IsActive: user.IsActive,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"), Posts: count,
	}
	return &resp, nil
}

// Delete removes a user by ID.
func (s *UserService) Delete(id uint) error {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}
	return s.DB.Delete(&user).Error
}