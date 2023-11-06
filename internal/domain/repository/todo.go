package repository

import (
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
)

type Todo interface {
	GetByUserID(userId uint64) ([]*model.Todo, error)
	Create(todo *model.Todo) (*model.Todo, error)
}
