package user_service

import (
	"kanban_board/dto"
	"kanban_board/pkg/errs"
)

type UserService interface {
	CreateNewUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
}
