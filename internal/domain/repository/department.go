package repository

import (
	"context"

	"github.com/hareta0109/graphql_sandbox/internal/domain/model/graph"
)

type Department interface {
	Get(id uint64) (*graph.Department, error)
	BulkGet() ([]*graph.Department, error)
	Create(department *graph.Department) (*graph.Department, error)

	GetMapInIDs(ctx context.Context, ids []uint64) (map[uint64]*graph.Department, error)
}
