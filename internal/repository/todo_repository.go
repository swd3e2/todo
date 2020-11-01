package repository

import "github.com/swd3e2/todo/internal/domain"

type TodoRepository interface {
	FindById(id int) *domain.Todo
}
