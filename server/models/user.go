package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"              json:"id"`
	Username  string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"json:"email"`
	FullName  string    `gorm:"type:varchar(100);not null"            json:"full_name"`
	Password  string    `gorm:"type:varchar(255);not null"            json:"-"`
	Role      string    `gorm:"type:varchar(20);default:'member'"     json:"role"`
	IsActive  bool      `gorm:"default:true"                          json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime"                        json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updated_at"`
}