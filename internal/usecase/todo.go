package usecase

import (
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
	"github.com/hareta0109/graphql_sandbox/internal/domain/repository"
)

type Todo interface {
	GetByUserID(userID uint64) ([]*model.Todo, error)
	Create(text string, userID uint64) (*model.Todo, error)
}

type TodoUsecase struct {
	todoRepo repository.Todo
}

func NewTodoUsecase(todoRepo repository.Todo) Todo {
	todoUsecase := TodoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

func (u *TodoUsecase) GetByUserID(userID uint64) ([]*model.Todo, error) {
	todo, err := u.todoRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *TodoUsecase) Create(text string, userID uint64) (*model.Todo, error) {
	todo := model.NewTodo(
		model.TodoDescription{
			Text:   text,
			UserID: userID,
		},
	)

	created, err := u.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return created, nil
}
