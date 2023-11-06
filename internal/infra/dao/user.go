package dao

import (
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
)

type User struct {
	ID           uint64
	DepartmentID uint64
	Name         string
}

func NewUser(user *model.User) *User {
	if user == nil {
		return nil
	}

	return &User{
		DepartmentID: user.DepartmentID,
		Name:         user.Name,
	}
}

func (d *User) ToModel() *model.User {
	if d == nil {
		return nil
	}

	return &model.User{
		ID:           d.ID,
		DepartmentID: d.DepartmentID,
		Name:         d.Name,
	}
}
