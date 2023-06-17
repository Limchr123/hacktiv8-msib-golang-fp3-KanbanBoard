package dto

import "time"

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserTask struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"User"`
}

type NewTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateStatusRequest struct {
	Status bool `json:"status"`
}

type UpdateTaskCategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type NewTask struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetAllTask struct {
	Status int        `json:"status"`
	Data   []UserTask `json:"data"`
}

type UpdateTask struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewTaskResponse struct {
	Status int     `json:"status"`
	Data   NewTask `json:"data"`
}

type UpdateTaskResponse struct {
	Status int        `json:"status"`
	Data   UpdateTask `json:"data"`
}

type DeleteTask struct {
	Message string `json:"message"`
}

type DeleteTaskResponse struct {
	Status int        `json:"status"`
	Data   DeleteTask `json:"data"`
}
