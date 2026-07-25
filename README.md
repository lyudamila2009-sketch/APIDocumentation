# Task Management API

REST API для управления пользователями и задачами на Go.

## Описание

Проект реализует простой REST API сервис с возможностями:
- Управление пользователями (CRUD операции)
- Управление задачами (CRUD операции)
- Полная документация API через Swagger UI

## Структура проекта

```
api-doc-example/
├── docs/                 # Документация API (Swagger)
│   ├── swagger.json     # OpenAPI спецификация
│   └── README.md        # Документация в docs/
├── internal/
│   ├── handlers/        # Обработчики HTTP запросов
│   │   ├── user.go      # Handler для управления пользователями
│   │   └── todo.go      # Handler для управления задачами
│   └── models/          # Модели данных
│       ├── user.go      # Модели пользователей
│       └── todo.go      # Модели задач
├── main.go              # Точка входа в приложение
├── go.mod               # Зависимости Go
├── go.sum               # Хэши зависимостей
└── README.md            # Документация проекта
```

> **Примечание:** Это репозиторий `APIDocumentation` для проекта `api-doc-example`, реализующего REST API сервис **Task Management API**.

## Установка и запуск

### Требования

- Go 1.22+

### Установка зависимостей

```bash
go mod tidy
```

### Запуск сервера

```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`

## Зависимости

### Основные

- `github.com/go-chi/chi/v5` - HTTP роутер
- `github.com/swaggo/http-swagger/v2` - Swagger UI
- `github.com/swaggo/swag` - генератор Swagger документации

### Установка swag (для генерации документации)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Генерация документации:

```bash
swag init -g main.go -o docs
```

## API Документация

После запуска сервера доступна документация API:

- **Swagger UI**: http://localhost:8080/swagger/
- **OpenAPI JSON**: http://localhost:8080/docs/swagger.json

### Эндпоинты

#### Пользователи (`/api/v1/users`)

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/api/v1/users` | Список всех пользователей |
| GET | `/api/v1/users/{id}` | Получить пользователя по ID |
| POST | `/api/v1/users` | Создать нового пользователя |
| PUT | `/api/v1/users/{id}` | Обновить пользователя |
| DELETE | `/api/v1/users/{id}` | Удалить пользователя |

#### Задачи (`/api/v1/todos`)

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| GET | `/api/v1/todos` | Список всех задач |
| GET | `/api/v1/todos/{id}` | Получить задачу по ID |
| POST | `/api/v1/todos` | Создать новую задачу |
| PUT | `/api/v1/todos/{id}` | Обновить задачу |
| DELETE | `/api/v1/todos/{id}` | Удалить задачу |

## Примеры запросов и ответов

### Создание пользователя

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe"
  }'
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

### Список задач

**Request:**
```bash
curl http://localhost:8080/api/v1/todos
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "title": "Learn Go",
      "description": "Learn Go for backend development",
      "done": false,
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ]
}
```

### Создание задачи

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "title": "Learn Go",
    "description": "Learn Go for backend development",
    "done": false
  }'
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Learn Go",
    "description": "Learn Go for backend development",
    "done": false,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

## Модели данных

В проекте используются Go структуры для описания данных. Ниже приведены описания моделей для справки.

### User

```go
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
```

### CreateUserRequest

```go
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,max=100"`
}
```

### UpdateUserRequest

```go
type UpdateUserRequest struct {
	Email *string `json:"email" validate:"omitempty,email"`
	Name  *string `json:"name" validate:"omitempty,max=100"`
}
```

### Todo

```go
type Todo struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
```

### CreateTodoRequest

```go
type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description" validate:"max=1000"`
	Done        bool   `json:"done"`
}
```

### UpdateTodoRequest

```go
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Done        *bool   `json:"done"`
}
```

### ErrorResponse

```go
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
```

### SuccessResponse

```go
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}
```

## Лицензия

MIT

## Контакты

Для вопросов и предложений создайте issue в репозитории проекта.
