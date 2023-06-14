package http_user

import (
	"github.com/gin-gonic/gin"
	"kanban_board/dto"
	"kanban_board/handler/user_handler"
	"kanban_board/pkg/errs"
	"kanban_board/service/user_service"
)

type userHandler struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) user_handler.UserHandler {
	return &userHandler{userService: userService}
}

func (u *userHandler) CreateNewUser(ctx *gin.Context) {
	var newUser dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		errBindJson := errs.NewUnproccesibleEntity("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := u.userService.CreateNewUser(&newUser)
	if err != nil {
		errBindJson := errs.NewBadRequest("Error occurred because request body is invalid")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(result.Status, result)
}
