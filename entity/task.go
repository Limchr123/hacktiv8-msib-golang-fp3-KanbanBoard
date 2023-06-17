package entity

import (
	"gorm.io/gorm"
	"kanban_board/pkg/errs"
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description;not null"`
	Status      bool      `gorm:"column:status;not null"`
	UserID      uint      `gorm:"column:user_id;"`
	User        User      `gorm:"foreignKey:UserID"`
	CategoryID  uint      `gorm:"column:category_id;"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `gorm:"column:created_at;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Category{}).Where("id = ?", t.CategoryID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errs.NewInternalServerError("Category not found")
	}

	return nil
}
