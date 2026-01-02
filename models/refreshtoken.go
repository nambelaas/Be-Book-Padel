package models

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `json:"user_id"`
	Token     string    `gorm:"text;not null" json:"token"`
	ExpiresAt time.Time `gorm:"timestamptz" json:"expires_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User      Users     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
}
