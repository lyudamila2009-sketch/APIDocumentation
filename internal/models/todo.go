package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description" validate:"max=1000"`
	Done        bool   `json:"done"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Done        *bool   `json:"done"`
}
