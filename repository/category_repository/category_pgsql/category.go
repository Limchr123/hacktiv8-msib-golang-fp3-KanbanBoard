package category_pgsql

import (
	"gorm.io/gorm"
	"kanban_board/entity"
	"kanban_board/pkg/errs"
	"kanban_board/repository/category_repository"
)

type categoryPg struct {
	db *gorm.DB
}

func NewCategoryPG(database *gorm.DB) category_repository.CategoryRepository {
	return &categoryPg{db: database}
}

func (c *categoryPg) CreateNewCategory(categoryPayload *entity.Category) (*entity.Category, errs.MessageErr) {
	if err := c.db.Create(categoryPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create category")
	}

	return categoryPayload, nil
}
