package usecase

type UserColumns struct {
	ID    int
	Name  string
	Todos []TodoColumns
}

type UserGetter interface {
	GetUsers() ([]UserColumns, error)
}

type User struct {
	userGetter UserGetter
	todoGetter TodoGetter
}

func NewUser(userGetter UserGetter, todoGetter TodoGetter) *User {
	return &User{userGetter, todoGetter}
}

func (u *User) GetUsers() ([]UserColumns, error) {
	us, err := u.userGetter.GetUsers()
	if err != nil {
		return nil, err
	}
	for i, user := range us {
		ts, err := u.todoGetter.GetTodosByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		us[i].Todos = ts
	}
	return us, err
}
