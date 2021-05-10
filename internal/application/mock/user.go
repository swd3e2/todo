package mock

import (
	"errors"
	"github.com/swd3e2/todo/internal/application"
)

var (
	userNotFound = errors.New("user not found")
)

// UserRepository Репозиторий пользователей
type UserRepository struct {
	db map[string]*application.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByLogin(login string) (*application.User, error) {
	if u, ok := r.db[login]; !ok {
		return nil, userNotFound
	} else {
		return u, nil
	}
}

func (r *UserRepository) Save(user *application.User) error {
	r.db[user.Login] = user
	return nil
}
