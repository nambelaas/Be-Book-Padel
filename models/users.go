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
	FirstName  string    `gorm:"type:varchar(100);not null;default:''" json:"first_name"`
	LastName   string    `gorm:"type:varchar(100);not null;default:''" json:"last_name"`
	Email      string    `gorm:"type:varchar(100);index:unique;not null;default:''" json:"email"`
	Password   string    `gorm:"type:varchar(100);not null;default:''" json:"password"`
	Gender     string    `gorm:"type:varchar(20);not null;default:''" json:"gender"`
	Dob        time.Time `gorm:"type:date" json:"dob"`
	Role       UserRole  `gorm:"type:user_roles;not null;default:'user'" json:"role"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
