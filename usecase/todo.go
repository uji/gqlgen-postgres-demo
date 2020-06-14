package usecase

type TodoColumns struct {
	ID   int
	Text string
	Done bool
	Name string
}

type TodoGetter interface {
	GetTodosByUserID(userID int) ([]TodoColumns, error)
}

type Todo struct {
	getter TodoGetter
}

func NewTodo(getter TodoGetter) *Todo {
	return &Todo{getter}
}

func (u *Todo) GetTodosByUserID(userID int) ([]TodoColumns, error) {
	return u.getter.GetTodosByUserID(userID)
}
