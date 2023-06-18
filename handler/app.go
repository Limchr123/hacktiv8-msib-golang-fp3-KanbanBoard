package handler

import (
	"github.com/gin-gonic/gin"
	"kanban_board/database"
	"kanban_board/handler/category_handler/http_category"
	"kanban_board/handler/task_handler/http_task"
	"kanban_board/handler/user_handler/http_user"
	"kanban_board/repository/category_repository/category_pgsql"
	"kanban_board/repository/task_repository/task_pgsql"
	"kanban_board/repository/user_repository/user_pgsql"
	"kanban_board/service/category_service/category"
	"kanban_board/service/task_service/task"
	"kanban_board/service/user_service/user"
	"os"
)

func StartApp() {
	var port = os.Getenv("PORT")
	db := database.GetDataBaseInstance()

	userRepo := user_pgsql.NewUserPG(db)
	userService := user.NewUserService(userRepo)
	userHandler := http_user.NewUserHandler(userService)

	categoryRepo := category_pgsql.NewCategoryPG(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := http_category.NewCategoryHandler(categoryService)

	taskRepo := task_pgsql.NewTaskPG(db)
	taskService := task.NewTaskService(taskRepo)
	taskHandler := http_task.NewTaskHandler(taskService)

	//Route
	r := gin.Default()

	r.POST("/users/register", userHandler.CreateNewUser)
	r.POST("/users/login", userHandler.UserLogin)
	userRoute := r.Group("/users")
	{
		userRoute.Use(userService.UserAuthentication())
		userRoute.POST("/update-account", userService.UserAuthorization(), userHandler.UpdateUserData)
		userRoute.DELETE("/delete-account", userService.UserAuthentication(), userHandler.DeleteUserData)
	}

	categoryRoute := r.Group("/categories")
	{
		categoryRoute.Use(userService.UserAuthentication())
		categoryRoute.POST("", userService.AdminAuthorization(), categoryHandler.CreateNewCategory)
		categoryRoute.GET("", userService.UserAuthorization(), categoryHandler.GetAllTaskByCategory)
		categoryRoute.PATCH("/:categoryId", userService.AdminAuthorization(), categoryHandler.UpdateCategoryById)
		categoryRoute.DELETE("/:categoryId", userService.AdminAuthorization(), categoryHandler.DeleteCategoryById)
	}

	taskRoute := r.Group("/tasks")
	{
		taskRoute.Use(userService.UserAuthentication())
		taskRoute.POST("", userService.UserAuthorization(), taskHandler.CreateNewTask)
		taskRoute.GET("", userService.UserAuthorization(), taskHandler.GetAllTask)
		taskRoute.PUT("/:taskId", userService.UserAuthorization(), taskHandler.UpdateTaskById)
		taskRoute.PATCH("/update-status/:taskId", userService.UserAuthorization(), taskHandler.UpdateStatusTask)
		taskRoute.PATCH("/update-category/:taskId", userService.UserAuthorization(), taskHandler.UpdateTaskCategory)
		taskRoute.DELETE("/:taskId", userService.UserAuthorization(), taskHandler.DeleteTaskById)
	}

	r.Run(":" + port)
}
