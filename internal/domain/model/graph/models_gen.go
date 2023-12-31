// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewDepartment struct {
	Name string `json:"name"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name         string `json:"name"`
	DepartmentID string `json:"departmentId"`
}
