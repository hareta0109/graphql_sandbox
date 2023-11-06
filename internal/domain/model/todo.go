package model

type Todo struct {
	ID     uint64 `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID uint64 `json:"userId"`
}

type TodoDescription struct {
	Text   string
	UserID uint64
}

func NewTodo(desc TodoDescription) *Todo {
	return &Todo{
		Text:   desc.Text,
		Done:   false,
		UserID: desc.UserID,
	}
}
