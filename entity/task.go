package entity

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description;not null"`
	Status      bool      `gorm:"column:status;not null"`
	UserID      int       `gorm:"column:user_id;"`
	User        User      `gorm:"foreignKey:UserID"`
	CategoryID  int       `gorm:"column:category_id;"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `gorm:"column:created_at;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;"`
}
