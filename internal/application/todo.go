package application

//go:generate mockgen -source=todo.go -destination=generated_mocks/todo_mock.go

// Todo Объект задачи
type Todo struct {
	UserId int
	Name   string
	Desc   string
}

// TodoRepository Репозиторий задач
type TodoRepository interface {
	FindByUserId(id int) *Todo
}
