package repository

import (
	mysql "github.com/dao0203/todo-sample/internal/database"
	"github.com/dao0203/todo-sample/internal/entity"
)

// https://github.com/evrone/go-clean-template/blob/master/internal/repo/persistent/translation_postgres.go
// を参考に実装

type TodoRepository struct {
	*mysql.MySql
}

func New(ms *mysql.MySql) *TodoRepository {
	return &TodoRepository{ms}
}

func (tr *TodoRepository) GetTodos() ([]entity.Todo, error) {
	rows, err := tr.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CategoryId, &todo.IsCompleted, &todo.DueDate); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
