package repository

import (
	"context"

	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
)

type User interface {
	Get(id uint64) (*model.User, error)
	BulkGet() ([]*model.User, error)
	Create(user *model.User) (*model.User, error)

	GetMapInIDs(ctx context.Context, ids []uint64) (map[uint64]*model.User, error)
}
