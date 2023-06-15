package category

import (
	"kanban_board/dto"
	"kanban_board/entity"
	"kanban_board/pkg/errs"
	"kanban_board/pkg/helpers"
	"kanban_board/repository/category_repository"
	"kanban_board/service/category_service"
	"net/http"
)

type categoryService struct {
	categoryRepo category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) category_service.CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (c *categoryService) CreateNewCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	category := entity.Category{
		Type: payload.Type,
	}

	result, err := c.categoryRepo.CreateNewCategory(&category)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create payload")
	}

	response := &dto.NewCategoryResponse{
		Status: http.StatusCreated,
		Data: dto.CategoryResponse{
			ID:        result.ID,
			Type:      result.Type,
			CreatedAt: result.CreatedAt,
		},
	}

	return response, nil
}

func (c *categoryService) UpdateCategoryById(id uint, payload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	category := entity.Category{
		Type: payload.Type,
	}

	result, err := c.categoryRepo.UpdateCategoryById(id, &category)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create payload")
	}

	response := &dto.UpdateCategoryResponse{
		Status: http.StatusOK,
		Data: dto.UpdateResponse{
			ID:        result.ID,
			Type:      result.Type,
			UpdatedAt: result.UpdatedAt,
		},
	}

	return response, nil
}

func (c *categoryService) DeleteCategoryById(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	if err := c.categoryRepo.DeleteCategoryById(id); err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to delete data")
	}

	response := &dto.DeleteCategoryResponse{
		Status:  http.StatusOK,
		Message: "Category has been successfully deleted",
	}

	return response, nil
}