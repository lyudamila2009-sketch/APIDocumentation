package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"

	"api-doc-example/internal/models"
	"api-doc-example/internal/validation"
)

type UserHandler struct {
	mu     sync.RWMutex
	users  map[int]models.User
	nextID int
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

// ListUsers godoc
// @Summary      Список пользователей
// @Description  Возвращает всех пользователей
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.SuccessResponse{data=[]models.User}
// @Router       /api/v1/users [get]
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	users := make([]models.User, 0, len(h.users))
	for _, user := range h.users {
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    users,
	})
}

// GetUser godoc
// @Summary      Получить пользователя
// @Description  Возвращает пользователя по ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID пользователя"
// @Success      200  {object}  models.SuccessResponse{data=models.User}
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	h.mu.RLock()
	user, exists := h.users[id]
	h.mu.RUnlock()

	if !exists {
		writeError(w, http.StatusNotFound, "Пользователь не найден")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    user,
	})
}

// CreateUser godoc
// @Summary      Создать пользователя
// @Description  Создаёт нового пользователя
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      models.CreateUserRequest  true  "Данные пользователя"
// @Success      201   {object}  models.SuccessResponse{data=models.User}
// @Failure      400   {object}  models.ErrorResponse
// @Router       /api/v1/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	if err := validation.Validate(req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	h.mu.Lock()
	user := models.User{
		ID:        h.nextID,
		Email:     req.Email,
		Name:      req.Name,
		CreatedAt: time.Now(),
	}
	h.users[h.nextID] = user
	h.nextID++
	h.mu.Unlock()

	log.Printf("Создан новый пользователь: %+v", user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    user,
	})
}

// UpdateUser godoc
// @Summary      Частично обновить пользователя
// @Description  Частично обновляет данные пользователя по ID. Поля с null остаются без изменений.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int                       true  "ID пользователя"
// @Param        body  body      models.UpdateUserRequest  true  "Обновляемые поля"
// @Success      200   {object}  models.SuccessResponse{data=models.User}
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/v1/users/{id} [patch]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	if err := validation.Validate(req); err != nil {
		writeError(w, http.StatusBadRequest, "Невалидный запрос")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	user, exists := h.users[id]
	if !exists {
		writeError(w, http.StatusNotFound, "Пользователь не найден")
		return
	}

	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Name != nil {
		user.Name = *req.Name
	}

	h.users[id] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    user,
	})
}

// DeleteUser godoc
// @Summary      Удалить пользователя
// @Description  Удаляет пользователя по ID
// @Tags         users
// @Param        id  path      int  true  "ID пользователя"
// @Success      204
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.users[id]; !exists {
		writeError(w, http.StatusNotFound, "Пользователь не найден")
		return
	}

	delete(h.users, id)

	w.WriteHeader(http.StatusNoContent)
}

// UserExists проверяет, существует ли пользователь с заданным ID
func (h *UserHandler) UserExists(id int) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	_, exists := h.users[id]
	return exists
}
