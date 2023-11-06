package model

import "github.com/hareta0109/graphql_sandbox/internal/domain/model/graph"

type DepartmentDescription struct {
	Name string
}

func NewDepartment(desc DepartmentDescription) *graph.Department {
	return &graph.Department{
		Name: desc.Name,
	}
}
