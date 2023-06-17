package category_repository

import (
	"kanban_board/entity"
	"kanban_board/pkg/errs"
)

type CategoryRepository interface {
	CreateNewCategory(categoryPayload *entity.Category) (*entity.Category, errs.MessageErr)
	GetCategoryById(id uint) (*entity.Category, errs.MessageErr)
	GetTaskByCategories() ([]entity.Category, errs.MessageErr)
	UpdateCategoryById(id uint, categoryPayload *entity.Category) (*entity.Category, errs.MessageErr)
	DeleteCategoryById(id uint) errs.MessageErr
}
