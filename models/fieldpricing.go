package models

import "time"

type DayType string

const (
	Weekday DayType = "weekday"
	Weekend DayType = "weekend"
)

type FieldPricing struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FieldID         uint      `gorm:"not null" json:"field_id"`
	Price           float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	DurationMinutes int       `gorm:"type:int;not null" json:"duration_minutes"`
	StartTime       time.Time `gorm:"type:datetime;not null" json:"start_time"`
	EndTime         time.Time `gorm:"type:datetime;not null" json:"end_time"`
	DayType         DayType   `gorm:"type:day_type;not null;default:'weekday'" json:"day_type"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
