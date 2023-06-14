package handler

import (
	"github.com/gin-gonic/gin"
	"kanban_board/database"
	"kanban_board/handler/user_handler/http_user"
	"kanban_board/repository/user_repository/user_pgsql"
	"kanban_board/service/user_service/user"
)

func StartApp() {
	db := database.GetDataBaseInstance()

	userRepo := user_pgsql.NewUserPG(db)
	userService := user.NewUserService(userRepo)
	userHandler := http_user.NewUserHandler(userService)

	//Route
	r := gin.Default()

	r.POST("/users/register", userHandler.CreateNewUser)
	r.POST("/users/login", userHandler.UserLogin)
	userRoute := r.Group("/users")
	{
		userRoute.Use(userService.UserAuthentication())
		userRoute.POST("/update-account", userService.UserAuthorization(), userHandler.UpdateUserData)
	}

	r.Run(":8080")
}
