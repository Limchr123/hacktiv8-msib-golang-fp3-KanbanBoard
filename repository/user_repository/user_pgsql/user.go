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

func (u *userPg) GetUserById(id uint) (*entity.User, errs.MessageErr) {
	var user entity.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find user")
	}

	return &user, nil
}

func (u *userPg) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User
	if err := u.db.Where("email = ?", userEmail).First(&user).Error; err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to find email")
	}

	return &user, nil
}

func (u *userPg) UpdateUserById(id uint, userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user, err := u.GetUserById(id)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find data")
	}

	if err := u.db.Model(user).Updates(userPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to update data")
	}

	return user, nil
}

func (u *userPg) DeleteUserById(id uint) errs.MessageErr {
	user, err := u.GetUserById(id)
	if err != nil {
		return errs.NewInternalServerError("Error occurred while trying to find data")
	}

	if err := u.db.Delete(user).Error; err != nil {
		return errs.NewInternalServerError("Error occurred while trying to delete data")
	}

	return nil
}
