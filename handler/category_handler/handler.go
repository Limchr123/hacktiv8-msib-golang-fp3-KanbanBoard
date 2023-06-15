package category_handler

import "github.com/gin-gonic/gin"

type CategoryHandler interface {
	CreateNewCategory(ctx *gin.Context)
}
