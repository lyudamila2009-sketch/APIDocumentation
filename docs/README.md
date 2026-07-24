# Task Management API Documentation

Swagger/OpenAPI документация для REST API сервиса управления задачами.

## Обзор API

- **Title**: Task Management API
- **Version**: 1.0
- **Description**: REST API для управления пользователями и задачами
- **Host**: localhost:8080
- **Base Path**: /

## Доступ к документации

### Swagger UI

Откройте в браузере:
```
http://localhost:8080/swagger/
```

### OpenAPI JSON

Спецификация в формате OpenAPI 2.0 (Swagger):
```
http://localhost:8080/docs/swagger.json
```

## Эндпоинты

### Пользователи (users)

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/v1/users` | Список всех пользователей |
| GET | `/api/v1/users/{id}` | Получить пользователя по ID |
| POST | `/api/v1/users` | Создать нового пользователя |
| PUT | `/api/v1/users/{id}` | Обновить пользователя |
| DELETE | `/api/v1/users/{id}` | Удалить пользователя |

### Задачи (todos)

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/v1/todos` | Список всех задач |
| GET | `/api/v1/todos/{id}` | Получить задачу по ID |
| POST | `/api/v1/todos` | Создать новую задачу |
| PUT | `/api/v1/todos/{id}` | Обновить задачу |
| DELETE | `/api/v1/todos/{id}` | Удалить задачу |

## Модели данных

### User

Пользователь системы.

```json
{
  "id": 1,
  "email": "user@example.com",
  "name": "John Doe",
  "created_at": "2024-01-01T12:00:00Z"
}
```

### Todo

Задача пользователя.

```json
{
  "id": 1,
  "user_id": 1,
  "title": "Learn Go",
  "description": "Learn Go for backend development",
  "done": false,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### SuccessResponse

Успешный ответ от API.

```json
{
  "success": true,
  "data": { ... }
}
```

### ErrorResponse

Ответ об ошибке.

```json
{
  "error": "error_code",
  "message": "Описание ошибки"
}
```

## Запросы

### Создание пользователя

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe"
  }'
```

### Создание задачи

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

## Коды ответов

| Код | Описание |
|-----|----------|
| 200 | Успешный запрос |
| 201 | Ресурс создан |
| 204 | Успешное удаление (нет содержимого) |
| 400 | Невалидный запрос |
| 404 | Ресурс не найден |

## Генерация документации

Для обновления `swagger.json` после изменений в коде:

```bash
swag init
```

## Файлы

- `swagger.json` - Сгенерированная OpenAPI спецификация
- `README.md` - Документация проекта
