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

func (u *userPg) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User
	if err := u.db.Where("email = ?", userEmail).First(&user).Error; err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to find email")
	}

	return &user, nil
}
