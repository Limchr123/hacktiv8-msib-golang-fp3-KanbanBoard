package user_service

import (
	"github.com/gin-gonic/gin"
	"kanban_board/dto"
	"kanban_board/pkg/errs"
)

type UserService interface {
	CreateNewUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	UserLogin(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUserData(id uint, payload *dto.UpdateRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	UserAuthorization() gin.HandlerFunc
	CategoryAuthorization() gin.HandlerFunc
	UserAuthentication() gin.HandlerFunc
}
