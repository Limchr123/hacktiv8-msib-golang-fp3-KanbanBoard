package http_category

import (
	"github.com/gin-gonic/gin"
	"kanban_board/dto"
	"kanban_board/handler/category_handler"
	"kanban_board/pkg/errs"
	"kanban_board/service/category_service"
)

type categoryHandler struct {
	categoryService category_service.CategoryService
}

func NewCategoryHandler(categoryService category_service.CategoryService) category_handler.CategoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

func (c *categoryHandler) CreateNewCategory(ctx *gin.Context) {
	var categoryRequest dto.NewCategoryRequest

	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := c.categoryService.CreateNewCategory(&categoryRequest)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}
