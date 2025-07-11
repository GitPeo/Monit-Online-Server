package models

type TodoResp struct {
	Id      int     `json:"tid"`
	Title   string  `json:"title"`
	Checked bool    `json:"checked"`
	Order   float64 `json:"order"`
	Version int     `json:"version"`
}

func TodoToResponse(todo Todo) TodoResp {
	return TodoResp{
		Id:      todo.Tid,
		Title:   todo.Title,
		Checked: todo.Checked,
		Order:   todo.Order,
		Version: todo.Version,
	}
}

// TodosToResponses 将 []model.Todo 转换为 []TodoResponse
func TodosToResponses(todos []Todo) []TodoResp {
	responses := make([]TodoResp, len(todos))
	for i, todo := range todos {
		responses[i] = TodoToResponse(todo)
	}
	return responses
}
