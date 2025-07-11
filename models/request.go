package models

type TodoReq struct {
	Tid     int     `json:"tid"`
	Title   string  `json:"title"`
	Checked bool    `json:"checked"`
	Order   float64 `json:"order"`
	Version int     `json:"version"`
	Deleted bool    `json:"deleted"`
}

func RequestToTodo(req TodoReq) Todo {
	return Todo{
		Tid:     req.Tid,
		Title:   req.Title,
		Checked: req.Checked,
		Order:   req.Order,
		Version: req.Version,
		Deleted: req.Deleted,
	}
}

// TodosToResponses 将 []model.Todo 转换为 []TodoResponse
func RequestsToTodos(reqs []TodoReq) []Todo {
	todos := make([]Todo, len(reqs))
	for i, req := range reqs {
		todos[i] = RequestToTodo(req)
	}
	return todos
}
