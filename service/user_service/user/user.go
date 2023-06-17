package user

import (
	"github.com/gin-gonic/gin"
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

func (u *userService) UserLogin(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate struct")
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, errs.NewNotFound("Error occurred because email is invalid")
	}

	userPassword := user.ComparePassword(payload.Password)
	if !userPassword {
		return nil, errs.NewNotFound("Error occurred because password is invalid")
	}

	token := user.GenerateToken()

	response := &dto.LoginResponse{
		Status: http.StatusOK,
		Data:   dto.TokenResponse{Token: token},
	}
	return response, nil
}

func (u *userService) UpdateUserData(id uint, payload *dto.UpdateRequest) (*dto.UpdateUserResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, errs.NewUnproccesibleEntity("Error occurred while trying to validate struct")
	}

	updatedUser := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
	}

	user, err := u.userRepo.UpdateUserById(id, &updatedUser)
	if err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to find data")
	}

	response := &dto.UpdateUserResponse{
		Status: http.StatusOK,
		Data: dto.UpdateUser{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response, nil
}

func (u *userService) DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	if err := u.userRepo.DeleteUserById(id); err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to delete data")
	}

	response := &dto.DeleteUserResponse{
		Status: http.StatusOK,
		Data:   dto.DeleteUser{Message: "Your account has been successfully deleted"},
	}

	return response, nil
}

func (u *userService) UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		account, err := u.userRepo.GetUserById(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if user.ID != account.ID {
			errAuthorized := errs.NewUnauthorizedError("You are not authorized to edit data")
			ctx.AbortWithStatusJSON(errAuthorized.Status(), errAuthorized)
			return
		}

		ctx.Next()
	}
}

func (u *userService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		account, err := u.userRepo.GetUserById(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if account.Role != "admin" {
			errAuthorized := errs.NewUnauthorizedError("You are not authorized to edit data")
			ctx.AbortWithStatusJSON(errAuthorized.Status(), errAuthorized)
			return
		}

		ctx.Next()
	}
}

func (u *userService) UserAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := u.userRepo.GetUserByEmail(user.Email)
		if err != nil {
			errAuthentication := errs.NewUnathenticationError("Error occurred because token is invalid")
			ctx.AbortWithStatusJSON(errAuthentication.Status(), errAuthentication)
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
