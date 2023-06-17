package category_service

import (
	"kanban_board/dto"
	"kanban_board/pkg/errs"
)

type CategoryService interface {
	CreateNewCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	GetAllTaskByCategories() (*dto.GetAllTaskByCategoriesResponse, errs.MessageErr)
	UpdateCategoryById(id uint, payload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategoryById(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
}
