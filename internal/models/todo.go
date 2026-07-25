package models

import "time"

// Todo представляет задачу пользователя.
type Todo struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTodoRequest содержит данные для создания новой задачи.
type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description" validate:"max=1000"`
	Done        bool   `json:"done"`
}

// UpdateTodoRequest содержит данные для частичного обновления задачи.
// Все поля опциональны: nil означает «не изменять».
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Done        *bool   `json:"done"`
	UserID      *int    `json:"user_id"`
}
