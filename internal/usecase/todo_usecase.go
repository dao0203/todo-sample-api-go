package usecase

import (
	"fmt"

	"github.com/dao0203/todo-sample/internal/entity"
	"github.com/dao0203/todo-sample/internal/repository"
)

type TodoUsecase struct {
	*repository.TodoRepository
}

func New(
	tr *repository.TodoRepository,
) *TodoUsecase {
	return &TodoUsecase{tr}
}

func (tu *TodoUsecase) GetTodos() ([]entity.Todo, error) {
	result, err := tu.TodoRepository.GetTodos()

	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	return result, nil
}
