package user

import (
	"kanban_board/dto"
	"kanban_board/entity"
	"kanban_board/pkg/errs"
	"kanban_board/pkg/helpers"
	"kanban_board/repository/user_repository"
	"kanban_board/service/user_service"
	"net/http"
)

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) user_service.UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateNewUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate request")
	}

	user := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     "member",
	}

	if err := user.HashPassword(); err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to hash password")
	}

	if err := u.userRepo.CreateNewUser(&user); err != nil {
		return nil, errs.NewBadRequest("Error occurred while trying to create user")
	}

	response := &dto.NewUserResponse{
		Status: http.StatusCreated,
		Data: dto.NewUser{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}

	return response, nil
}
