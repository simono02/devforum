package models

import "time"

// Post represents a forum post written by a member.
// Only column definitions — no methods on this struct.
type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"          json:"id"`
	Title     string    `gorm:"type:varchar(300);not null"        json:"title"`
	Content   string    `gorm:"type:text;not null"                json:"content"`
	Tag       string    `gorm:"type:varchar(50)"                  json:"tag"`
	UserID    uint      `gorm:"not null;index"                    json:"user_id"`
	Upvotes   int       `gorm:"default:0"                         json:"upvotes"`
	IsPinned  bool      `gorm:"default:false"                     json:"is_pinned"`
	IsFlagged bool      `gorm:"default:false"                     json:"is_flagged"`
	CreatedAt time.Time `gorm:"autoCreateTime"                    json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                    json:"updated_at"`
}