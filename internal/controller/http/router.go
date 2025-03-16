package http

import (
	v1 "github.com/dao0203/todo-sample/internal/controller/http/v1"
	"github.com/dao0203/todo-sample/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

func NewRouter(app *fiber.App, us *usecase.TodoUsecase) {
	// Define the routes
	apiV1Group := app.Group("/api/v1")
	v1.NewTodoRoutes(us, apiV1Group)
}
