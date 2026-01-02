package models

import "time"

type UserRole string

const (
	Admin UserRole = "admin"
	User  UserRole = "user"
	Staff UserRole = "staff"
)

type Users struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName  string    `gorm:"text;not null" json:"first_name"`
	LastName   string    `gorm:"text;not null" json:"last_name"`
	Email      string    `gorm:"text;index:unique;not null" json:"email"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	Password   string    `gorm:"text;not null" json:"password"`
	Gender     string    `gorm:"text;not null" json:"gender"`
	Dob        time.Time `gorm:"type:date" json:"dob"`
	Role       UserRole  `gorm:"type:enum('admin','user','staff');default:'user'" json:"role"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
