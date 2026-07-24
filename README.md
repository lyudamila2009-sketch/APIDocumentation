# Task Management API

REST API для управления пользователями и задачами на Go.

## Описание

Проект реализует простой REST API сервис с возможностями:
- Управление пользователями (CRUD операции)
- Управление задачами (CRUD операции)
- Полная документация API через Swagger UI

## Структура проекта

```
ai-assist-it/
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

## Установка и запуск

### Требования

- Go 1.22+
- make (опционально)

### Установка зависимостей

```bash
go mod tidy
```

### Запуск сервера

```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`

### Запуск через make (если доступно)

```bash
make run
```

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
swag init
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

## Pydantic модели (описание)

> Примечание: В проекте используются Go структуры, но ниже приведено описание аналогов Pydantic моделей для понимания структуры данных.

### User

```python
class User(BaseModel):
    id: int
    email: str
    name: str
    created_at: datetime
```

### CreateUserRequest

```python
class CreateUserRequest(BaseModel):
    email: str = Field(..., pattern=r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$')
    name: str = Field(..., max_length=100)
```

### UpdateUserRequest

```python
class UpdateUserRequest(BaseModel):
    email: Optional[str] = Field(None, pattern=r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$')
    name: Optional[str] = Field(None, max_length=100)
```

### Todo

```python
class Todo(BaseModel):
    id: int
    user_id: int
    title: str
    description: str
    done: bool
    created_at: datetime
    updated_at: datetime
```

### CreateTodoRequest

```python
class CreateTodoRequest(BaseModel):
    user_id: int
    title: str = Field(..., max_length=200)
    description: Optional[str] = Field(None, max_length=1000)
    done: bool = False
```

### UpdateTodoRequest

```python
class UpdateTodoRequest(BaseModel):
    title: Optional[str] = Field(None, max_length=200)
    description: Optional[str] = Field(None, max_length=1000)
    done: Optional[bool] = None
```

### ErrorResponse

```python
class ErrorResponse(BaseModel):
    error: str
    message: str
```

### SuccessResponse

```python
class SuccessResponse(BaseModel):
    success: bool
    data: Optional[Any] = None
```

## Лицензия

MIT

## Контакты

Для вопросов и предложений создайте issue в репозитории проекта.
