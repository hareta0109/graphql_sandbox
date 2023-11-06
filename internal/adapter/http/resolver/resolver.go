package resolver

import "github.com/hareta0109/graphql_sandbox/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DepartmentUsecase usecase.Department
	UserUsecase       usecase.User
	TodoUsecase       usecase.Todo
}
