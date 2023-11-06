package infra

import (
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
	"github.com/hareta0109/graphql_sandbox/internal/domain/repository"
	"github.com/hareta0109/graphql_sandbox/internal/infra/dao"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.Todo {
	return &todoRepository{db}
}

func (r *todoRepository) GetByUserID(userId uint64) ([]*model.Todo, error) {
	var rows []dao.Todo
	if err := r.db.Where("user_id = ?", userId).Find(&rows).Error; err != nil {
		return nil, err
	}

	var todoes []*model.Todo
	for _, row := range rows {
		todo := row.ToModel()
		todoes = append(todoes, todo)
	}

	return todoes, nil
}

func (r *todoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	dbModel := dao.NewTodo(todo)
	if err := r.db.Create(dbModel).Error; err != nil {
		return nil, err
	}

	return dbModel.ToModel(), nil
}
