package models

import "time"

type Field struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string         `gorm:"type:varchar(100);not null" json:"name"`
	Length        int            `gorm:"type:int;not null" json:"length"`
	Width         int            `gorm:"type:int;not null" json:"width"`
	Description   string         `gorm:"type:text" json:"description"`
	IsActive      bool           `gorm:"type:boolean;default:true" json:"is_active"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	FieldPricings []FieldPricing `gorm:"foreignKey:FieldID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"field_pricings"`
}
