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

func (c *categoryPg) GetCategoryById(id uint) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	if err := c.db.First(&category, id).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find category")
	}

	return &category, nil
}

func (c *categoryPg) UpdateCategoryById(id uint, categoryPayload *entity.Category) (*entity.Category, errs.MessageErr) {
	category, err := c.GetCategoryById(id)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find category")
	}

	if err := c.db.Model(category).Updates(categoryPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to update category")
	}

	return category, nil
}
