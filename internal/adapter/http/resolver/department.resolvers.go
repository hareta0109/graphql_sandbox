package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/hareta0109/graphql_sandbox/internal/domain/model/graph"
)

// CreateDepartment is the resolver for the createDepartment field.
func (r *mutationResolver) CreateDepartment(ctx context.Context, input graph.NewDepartment) (*graph.Department, error) {
	panic(fmt.Errorf("not implemented: CreateDepartment - createDepartment"))
}

// Department is the resolver for the department field.
func (r *queryResolver) Department(ctx context.Context, id string) (*graph.Department, error) {
	panic(fmt.Errorf("not implemented: Department - department"))
}

// Departments is the resolver for the departments field.
func (r *queryResolver) Departments(ctx context.Context) ([]*graph.Department, error) {
	panic(fmt.Errorf("not implemented: Departments - departments"))
}
