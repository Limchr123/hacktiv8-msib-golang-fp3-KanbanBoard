package user_repository

import (
	"kanban_board/entity"
	"kanban_board/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(userPayload *entity.User) errs.MessageErr
	GetUserById(id uint) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	UpdateUserById(id uint, userPayload *entity.User) (*entity.User, errs.MessageErr)
}
