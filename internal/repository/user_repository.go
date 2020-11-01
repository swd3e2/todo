package repository

import "github.com/swd3e2/todo/internal/domain"

type UserRepository interface {
	FindByLogin(login string) (*domain.User, error)
	Save(user *domain.User) (*domain.User, error)
}
