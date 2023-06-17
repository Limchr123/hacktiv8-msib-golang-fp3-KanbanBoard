package task_repository

import (
	"kanban_board/entity"
	"kanban_board/pkg/errs"
)

type TaskRepository interface {
	CreateNewTask(taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	GetTaskById(id uint) (*entity.Task, errs.MessageErr)
	GetAllTask() ([]entity.Task, errs.MessageErr)
	UpdateTaskById(id uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	DeleteTaskById(id uint) errs.MessageErr
}
