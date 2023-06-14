package user_repository

import (
	"kanban_board/entity"
	"kanban_board/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(userPayload *entity.User) errs.MessageErr
}
