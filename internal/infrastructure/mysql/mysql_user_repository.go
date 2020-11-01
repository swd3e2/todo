package mysql

import (
	"database/sql"

	"github.com/swd3e2/todo/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (this *UserRepository) FindByLogin(login string) (user *domain.User, err error) {
	user = new(domain.User)

	if err = this.db.QueryRow(
		`select name, lastname, age, login, password from users where login = ?`,
		login,
	).Scan(&user.Name, &user.LastName, &user.Age, &user.Credentials.Login, &user.Credentials.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (this *UserRepository) Save(user *domain.User) (*domain.User, error) {
	if res, err := this.db.Exec(`
		insert into users (name, lastname, age, login, password) values (?, ?, ?, ?, ?)
	`,
		user.Name,
		user.LastName,
		user.Age,
		user.Credentials.Login,
		user.Credentials.Password,
	); err != nil {
		return nil, err
	} else {
		if id, err := res.LastInsertId(); err != nil {
			return nil, err
		} else {
			user.Id = id
		}
	}

	return user, nil
}
