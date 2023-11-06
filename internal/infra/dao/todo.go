package dao

import (
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
)

type Todo struct {
	ID     uint64
	UserID uint64
	Text   string
	Done   bool
}

func NewTodo(todo *model.Todo) *Todo {
	if todo == nil {
		return nil
	}

	return &Todo{
		UserID: todo.UserID,
		Text:   todo.Text,
		Done:   false,
	}
}

func (d *Todo) ToModel() *model.Todo {
	if d == nil {
		return nil
	}

	return &model.Todo{
		ID:     d.ID,
		UserID: d.UserID,
		Text:   d.Text,
		Done:   d.Done,
	}
}
