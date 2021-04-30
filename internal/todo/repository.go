package todo

type Repository interface {
	FindById(id int) *Todo
}
