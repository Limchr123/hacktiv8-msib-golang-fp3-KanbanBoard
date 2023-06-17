package task_handler

import "github.com/gin-gonic/gin"

type TaskHandler interface {
	CreateNewTask(ctx *gin.Context)
	UpdateTaskById(ctx *gin.Context)
	UpdateStatusTask(ctx *gin.Context)
	UpdateTaskCategory(ctx *gin.Context)
	GetAllTask(ctx *gin.Context)
	DeleteTaskById(ctx *gin.Context)
}
