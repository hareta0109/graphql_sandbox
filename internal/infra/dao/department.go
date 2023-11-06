package dao

import (
	"github.com/hareta0109/graphql_sandbox/internal/domain/model/graph"
)

type Department struct {
	ID   uint64
	Name string
}

func NewDepartment(department *graph.Department) *Department {
	if department == nil {
		return nil
	}

	return &Department{
		Name: department.Name,
	}
}

func (d *Department) ToModel() *graph.Department {
	if d == nil {
		return nil
	}

	return &graph.Department{
		ID:   d.ID,
		Name: d.Name,
	}
}
