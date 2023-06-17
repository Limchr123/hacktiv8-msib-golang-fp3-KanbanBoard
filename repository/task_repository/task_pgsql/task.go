package task_pgsql

import (
	"gorm.io/gorm"
	"kanban_board/entity"
	"kanban_board/pkg/errs"
	"kanban_board/repository/task_repository"
)

type taskPg struct {
	db *gorm.DB
}

func NewTaskPG(database *gorm.DB) task_repository.TaskRepository {
	return &taskPg{db: database}
}

func (t *taskPg) CreateNewTask(taskPayload *entity.Task) (*entity.Task, errs.MessageErr) {
	if err := t.db.Create(taskPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create task")
	}

	return taskPayload, nil
}

func (t *taskPg) GetTaskById(id uint) (*entity.Task, errs.MessageErr) {
	var task entity.Task
	if err := t.db.First(&task, id).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find data")
	}

	return &task, nil
}

func (t *taskPg) GetAllTask() ([]entity.Task, errs.MessageErr) {
	var allTask []entity.Task

	if err := t.db.Preload("User").Find(&allTask).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to get data")
	}

	return allTask, nil
}

func (t *taskPg) UpdateTaskById(id uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskById(id)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if err := t.db.Model(task).Updates(taskPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to update task")
	}

	return task, nil
}

func (t *taskPg) DeleteTaskById(id uint) errs.MessageErr {
	task, err := t.GetTaskById(id)
	if err != nil {
		return errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if err := t.db.Delete(task).Error; err != nil {
		return errs.NewInternalServerError("Error occurred while trying to update task")
	}

	return nil
}
