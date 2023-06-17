package http_task

import (
	"github.com/gin-gonic/gin"
	"kanban_board/dto"
	"kanban_board/entity"
	"kanban_board/handler/task_handler"
	"kanban_board/pkg/errs"
	"kanban_board/service/task_service"
	"strconv"
)

type taskHandler struct {
	taskService task_service.TaskService
}

func NewTaskHandler(taskService task_service.TaskService) task_handler.TaskHandler {
	return &taskHandler{taskService: taskService}
}

func (t *taskHandler) CreateNewTask(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	var newTask dto.NewTaskRequest
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := t.taskService.CreateNewTask(user.ID, &newTask)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (t *taskHandler) UpdateTaskById(ctx *gin.Context) {
	id := ctx.Param("taskId")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	var updateTask dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&updateTask); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := t.taskService.UpdateTaskById(uint(taskId), &updateTask)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (t *taskHandler) UpdateStatusTask(ctx *gin.Context) {
	id := ctx.Param("taskId")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	var updateStatus dto.UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&updateStatus); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := t.taskService.UpdateStatusTask(uint(taskId), &updateStatus)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (t *taskHandler) UpdateTaskCategory(ctx *gin.Context) {
	id := ctx.Param("taskId")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	var updateCategory dto.UpdateTaskCategoryRequest
	if err := ctx.ShouldBindJSON(&updateCategory); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := t.taskService.UpdateTaskCategory(uint(taskId), &updateCategory)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (t *taskHandler) GetAllTask(ctx *gin.Context) {
	result, err := t.taskService.GetAllTask()
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}

func (t *taskHandler) DeleteTaskById(ctx *gin.Context) {
	id := ctx.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because id not found")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := t.taskService.DeleteTaskById(uint(taskId))
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}
