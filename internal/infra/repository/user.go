package infra

import (
	"context"

	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
	"github.com/hareta0109/graphql_sandbox/internal/domain/repository"
	"github.com/hareta0109/graphql_sandbox/internal/infra/dao"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.User {
	return &userRepository{db}
}

func (r *userRepository) Get(id uint64) (*model.User, error) {
	var row dao.User
	if err := r.db.Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}

	user := row.ToModel()

	return user, nil
}

func (r *userRepository) BulkGet() ([]*model.User, error) {
	var rows []dao.User
	if err := r.db.Find(&rows).Error; err != nil {
		return nil, err
	}

	var users []*model.User
	for _, row := range rows {
		user := row.ToModel()
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	dbModel := dao.NewUser(user)
	if err := r.db.Create(dbModel).Error; err != nil {
		return nil, err
	}

	return dbModel.ToModel(), nil
}

func (r *userRepository) GetMapInIDs(ctx context.Context, ids []uint64) (map[uint64]*model.User, error) {
	var rows []dao.User
	if err := r.db.Find(&rows, ids).Error; err != nil {
		return nil, err
	}

	results := make(map[uint64]*model.User)
	for _, row := range rows {
		user := row.ToModel()
		results[user.ID] = user
	}

	return results, nil
}
