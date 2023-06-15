package handler

import (
	"github.com/gin-gonic/gin"
	"kanban_board/database"
	"kanban_board/handler/category_handler/http_category"
	"kanban_board/handler/user_handler/http_user"
	"kanban_board/repository/category_repository/category_pgsql"
	"kanban_board/repository/user_repository/user_pgsql"
	"kanban_board/service/category_service/category"
	"kanban_board/service/user_service/user"
)

func StartApp() {
	db := database.GetDataBaseInstance()

	userRepo := user_pgsql.NewUserPG(db)
	userService := user.NewUserService(userRepo)
	userHandler := http_user.NewUserHandler(userService)

	categoryRepo := category_pgsql.NewCategoryPG(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := http_category.NewCategoryHandler(categoryService)

	//Route
	r := gin.Default()

	r.POST("/users/register", userHandler.CreateNewUser)
	r.POST("/users/login", userHandler.UserLogin)
	userRoute := r.Group("/users")
	{
		userRoute.Use(userService.UserAuthentication())
		userRoute.POST("/update-account", userService.UserAuthorization(), userHandler.UpdateUserData)
	}

	categoryRoute := r.Group("/categories")
	{
		categoryRoute.Use(userService.UserAuthentication())
		categoryRoute.POST("", userService.CategoryAuthorization(), categoryHandler.CreateNewCategory)
	}

	r.Run(":8080")
}
