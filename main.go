// @title           Task Management API
// @version         1.0
// @description     REST API для управления пользователями и задачами
// @host            localhost:8080
// @BasePath        /
package main

import (
	"log"
	"net/http"

	"api-doc-example/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	swaggerDir := "docs"

	r.Mount("/swagger", httpSwagger.Handler())

	usersRouter := chi.NewRouter()
	userHandler := handlers.NewUserHandler()
	usersRouter.Get("/", userHandler.ListUsers)
	usersRouter.Get("/{id:[0-9]+}", userHandler.GetUser)
	usersRouter.Post("/", userHandler.CreateUser)
	usersRouter.Put("/{id:[0-9]+}", userHandler.UpdateUser)
	usersRouter.Delete("/{id:[0-9]+}", userHandler.DeleteUser)
	r.Mount("/api/v1/users", usersRouter)

	todosRouter := chi.NewRouter()
	todoHandler := handlers.NewTodoHandler()
	todoHandler.SetUserHandler(userHandler)
	todosRouter.Get("/", todoHandler.ListTodos)
	todosRouter.Get("/{id:[0-9]+}", todoHandler.GetTodo)
	todosRouter.Post("/", todoHandler.CreateTodo)
	todosRouter.Put("/{id:[0-9]+}", todoHandler.UpdateTodo)
	todosRouter.Delete("/{id:[0-9]+}", todoHandler.DeleteTodo)
	r.Mount("/api/v1/todos", todosRouter)

	filesServer := http.FileServer(http.Dir(swaggerDir))
	r.Handle("/docs/*", http.StripPrefix("/docs/", filesServer))

	log.Println("Сервер запущен на http://localhost:8080")
	log.Println("Swagger UI доступен по адресу: http://localhost:8080/swagger/")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}