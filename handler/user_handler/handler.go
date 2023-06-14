package user_handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateNewUser(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
}
