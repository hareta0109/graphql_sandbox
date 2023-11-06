package model

type User struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	DepartmentID uint64 `json:"departmentId"`
}

type UserDescription struct {
	Name         string
	DepartmentID uint64
}

func NewUser(desc UserDescription) *User {
	return &User{
		Name:         desc.Name,
		DepartmentID: desc.DepartmentID,
	}
}
