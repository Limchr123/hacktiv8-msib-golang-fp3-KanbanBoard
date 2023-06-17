package task

import (
	"kanban_board/dto"
	"kanban_board/entity"
	"kanban_board/pkg/errs"
	"kanban_board/pkg/helpers"
	"kanban_board/repository/task_repository"
	"kanban_board/service/task_service"
	"net/http"
)

type taskService struct {
	taskRepo task_repository.TaskRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository) task_service.TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (t *taskService) CreateNewTask(id uint, payload *dto.NewTaskRequest) (*dto.NewTaskResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	task := entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      false,
		UserID:      id,
		CategoryID:  payload.CategoryID,
	}

	result, err := t.taskRepo.CreateNewTask(&task)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create payload")
	}

	response := &dto.NewTaskResponse{
		Status: http.StatusCreated,
		Data: dto.NewTask{
			ID:          result.ID,
			Title:       result.Title,
			Status:      result.Status,
			Description: result.Description,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			CreatedAt:   result.CreatedAt,
		},
	}

	return response, nil
}

func (t *taskService) UpdateTaskById(id uint, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	updatedTask := entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
	}

	result, err := t.taskRepo.UpdateTaskById(id, &updatedTask)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to edit payload")
	}

	response := &dto.UpdateTaskResponse{
		Status: http.StatusOK,
		Data: dto.UpdateTask{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Status:      result.Status,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			UpdatedAt:   result.UpdatedAt,
		},
	}

	return response, nil
}

func (t *taskService) UpdateStatusTask(id uint, payload *dto.UpdateStatusRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	updateStatus := entity.Task{
		Status: payload.Status,
	}

	result, err := t.taskRepo.UpdateTaskById(id, &updateStatus)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to edit payload")
	}

	response := &dto.UpdateTaskResponse{
		Status: http.StatusOK,
		Data: dto.UpdateTask{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Status:      result.Status,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			UpdatedAt:   result.UpdatedAt,
		},
	}

	return response, nil
}

func (t *taskService) UpdateTaskCategory(id uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	updateStatus := entity.Task{
		CategoryID: payload.CategoryID,
	}

	result, err := t.taskRepo.UpdateTaskById(id, &updateStatus)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to edit payload")
	}

	response := &dto.UpdateTaskResponse{
		Status: http.StatusOK,
		Data: dto.UpdateTask{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Status:      result.Status,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			UpdatedAt:   result.UpdatedAt,
		},
	}

	return response, nil
}

func (t *taskService) GetAllTask() (*dto.GetAllTask, errs.MessageErr) {
	allTask, err := t.taskRepo.GetAllTask()
	if err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to get data")
	}

	userTask := []dto.UserTask{}
	for _, eachTask := range allTask {
		user := dto.UserTask{
			ID:          eachTask.ID,
			Title:       eachTask.Title,
			Description: eachTask.Description,
			UserID:      eachTask.UserID,
			CategoryID:  eachTask.CategoryID,
			CreatedAt:   eachTask.CreatedAt,
			User: dto.User{
				ID:       eachTask.User.ID,
				Email:    eachTask.User.Email,
				FullName: eachTask.User.FullName,
			},
		}
		userTask = append(userTask, user)
	}

	response := &dto.GetAllTask{
		Status: http.StatusOK,
		Data:   userTask,
	}

	return response, nil
}

func (t *taskService) DeleteTaskById(id uint) (*dto.DeleteTaskResponse, errs.MessageErr) {
	if err := t.taskRepo.DeleteTaskById(id); err != nil {
		return nil, errs.NewInternalServerError("Cannot delete")
	}

	response := &dto.DeleteTaskResponse{
		Status: http.StatusOK,
		Data:   dto.DeleteTask{Message: "Task has been successfully deleted"},
	}

	return response, nil
}
