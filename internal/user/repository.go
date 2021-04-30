package user

type Repository interface {
	FindByLogin(login string) (*User, error)
	Save(user *User) (*User, error)
}
