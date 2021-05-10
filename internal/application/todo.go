package application

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
