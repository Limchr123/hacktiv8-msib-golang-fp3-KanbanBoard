package category_handler

import "github.com/gin-gonic/gin"

type CategoryHandler interface {
	CreateNewCategory(ctx *gin.Context)
	UpdateCategoryById(ctx *gin.Context)
	DeleteCategoryById(ctx *gin.Context)
}
