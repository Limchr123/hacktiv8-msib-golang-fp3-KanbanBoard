package dto

import "time"

type NewCategoryRequest struct {
	Type string `json:"type"`
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewCategoryResponse struct {
	Status int              `json:"status"`
	Data   CategoryResponse `json:"data"`
}

type UpdateCategoryResponse struct {
	Status int            `json:"status"`
	Data   UpdateResponse `json:"data"`
}
