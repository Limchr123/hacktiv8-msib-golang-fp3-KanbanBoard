package http_category

import (
	"github.com/gin-gonic/gin"
	"kanban_board/dto"
	"kanban_board/handler/category_handler"
	"kanban_board/pkg/errs"
	"kanban_board/service/category_service"
	"strconv"
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

func (c *categoryHandler) GetAllTaskByCategory(ctx *gin.Context) {
	result, err := c.categoryService.GetAllTaskByCategories()
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (c *categoryHandler) UpdateCategoryById(ctx *gin.Context) {
	id := ctx.Param("categoryId")

	categoryId, err := strconv.Atoi(id)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	var updateCategory dto.NewCategoryRequest
	if err := ctx.ShouldBindJSON(&updateCategory); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := c.categoryService.UpdateCategoryById(uint(categoryId), &updateCategory)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (c *categoryHandler) DeleteCategoryById(ctx *gin.Context) {
	id := ctx.Param("categoryId")

	categoryId, err := strconv.Atoi(id)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := c.categoryService.DeleteCategoryById(uint(categoryId))
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}
