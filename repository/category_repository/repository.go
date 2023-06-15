package category_repository

import (
	"kanban_board/entity"
	"kanban_board/pkg/errs"
)

type CategoryRepository interface {
	CreateNewCategory(categoryPayload *entity.Category) (*entity.Category, errs.MessageErr)
}
