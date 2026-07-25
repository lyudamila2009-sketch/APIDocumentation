package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"api-doc-example/internal/models"
	"api-doc-example/internal/validation"
)

type TodoHandler struct {
	todos  map[int]models.Todo
	nextID int
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

// ListTodos godoc
// @Summary      Список задач
// @Description  Возвращает все задачи
// @Tags         todos
// @Produce      json
// @Success      200  {object}  models.SuccessResponse{data=[]models.Todo}
// @Router       /api/v1/todos [get]
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos := make([]models.Todo, 0, len(h.todos))
	for _, todo := range h.todos {
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todos,
	})
}

// GetTodo godoc
// @Summary      Получить задачу
// @Description  Возвращает задачу по ID
// @Tags         todos
// @Produce      json
// @Param        id   path      int  true  "ID задачи"
// @Success      200  {object}  models.SuccessResponse{data=models.Todo}
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/v1/todos/{id} [get]
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := parseID(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	todo, exists := h.todos[id]
	if !exists {
		writeError(w, http.StatusNotFound, "Задача не найдена")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// CreateTodo godoc
// @Summary      Создать задачу
// @Description  Создаёт новую задачу
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        body  body      models.CreateTodoRequest  true  "Данные задачи"
// @Success      201   {object}  models.SuccessResponse{data=models.Todo}
// @Failure      400   {object}  models.ErrorResponse
// @Router       /api/v1/todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	if err := validation.Validate(req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	todo := models.Todo{
		ID:          h.nextID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	h.todos[h.nextID] = todo
	h.nextID++

	log.Printf("Создана новая задача: %+v", todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// UpdateTodo godoc
// @Summary      Обновить задачу
// @Description  Частично обновляет задачу по ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      int                       true  "ID задачи"
// @Param        body  body      models.UpdateTodoRequest  true  "Обновляемые поля"
// @Success      200   {object}  models.SuccessResponse{data=models.Todo}
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/v1/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := parseID(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	if err := validation.Validate(req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	todo, exists := h.todos[id]
	if !exists {
		writeError(w, http.StatusNotFound, "Задача не найдена")
		return
	}

	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Done != nil {
		todo.Done = *req.Done
	}
	todo.UpdatedAt = time.Now()

	h.todos[id] = todo

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// DeleteTodo godoc
// @Summary      Удалить задачу
// @Description  Удаляет задачу по ID
// @Tags         todos
// @Param        id  path      int  true  "ID задачи"
// @Success      204
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/v1/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := parseID(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	if _, exists := h.todos[id]; !exists {
		writeError(w, http.StatusNotFound, "Задача не найдена")
		return
	}

	delete(h.todos, id)

	w.WriteHeader(http.StatusNoContent)
}

// parseID parses ID from path parameter
func parseID(idStr string) (int, error) {
	if idStr == "" {
		return 0, &parseError{}
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, &parseError{}
	}
	return id, nil
}

type parseError struct{}

func (e *parseError) Error() string {
	return "invalid ID"
}
