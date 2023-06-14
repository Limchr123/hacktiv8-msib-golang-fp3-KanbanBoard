package user_pgsql

import (
	"gorm.io/gorm"
	"kanban_board/entity"
	"kanban_board/pkg/errs"
	"kanban_board/repository/user_repository"
)

type userPg struct {
	db *gorm.DB
}

func NewUserPG(database *gorm.DB) user_repository.UserRepository {
	return &userPg{db: database}
}

func (u *userPg) CreateNewUser(userPayload *entity.User) errs.MessageErr {
	if err := u.db.Create(userPayload).Error; err != nil {
		return errs.NewInternalServerError("Error occurred while trying to create user")
	}

	return nil
}
