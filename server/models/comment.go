package models

import "time"

// Comment represents a reply to a forum post.
// Only column definitions — no methods on this struct.
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"         json:"id"`
	Content   string    `gorm:"type:text;not null"               json:"content"`
	UserID    uint      `gorm:"not null;index"                   json:"user_id"`
	PostID    uint      `gorm:"not null;index"                   json:"post_id"`
	IsEdited  bool      `gorm:"default:false"                    json:"is_edited"`
	CreatedAt time.Time `gorm:"autoCreateTime"                   json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                   json:"updated_at"`
}