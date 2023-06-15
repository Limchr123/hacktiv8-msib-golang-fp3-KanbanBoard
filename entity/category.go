package entity

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Type      string    `gorm:"column:type;not null"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
	Tasks     []Task    `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
