package v1

import (
	"github.com/dao0203/todo-sample/internal/usecase"
	fiber "github.com/gofiber/fiber/v3/"
)

type TodoRoutes struct {
	us *usecase.TodoUsecase
}

func NewTodoRoutes(us *usecase.TodoUsecase, app *fiber.Router) {
	route := &TodoRoutes{us: us}

	todoGroup := app.Group("/todo")
	todoGroup.Get("/list", route.GetTodos)
}

func (td *TodoRoutes) GetTodos(ctx *fiber.Ctx) error {
	todos, err := td.us.GetTodos()
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(todos)
}
