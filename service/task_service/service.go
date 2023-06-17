package task_service

import (
	"kanban_board/dto"
	"kanban_board/pkg/errs"
)

type TaskService interface {
	CreateNewTask(id uint, payload *dto.NewTaskRequest) (*dto.NewTaskResponse, errs.MessageErr)
	GetAllTask() (*dto.GetAllTask, errs.MessageErr)
	UpdateTaskById(id uint, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateStatusTask(id uint, payload *dto.UpdateStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateTaskCategory(id uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	DeleteTaskById(id uint) (*dto.DeleteTaskResponse, errs.MessageErr)
}
