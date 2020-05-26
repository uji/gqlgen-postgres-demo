package usecase

type UserColumns struct {
	ID   int
	Name string
}

type UserGetter interface {
	GetUsers() ([]UserColumns, error)
}

type User struct {
	getter UserGetter
}

func NewUser(getter UserGetter) *User {
	return &User{getter}
}

func (u *User) GetUsers() ([]UserColumns, error) {
	return u.getter.GetUsers()
}
