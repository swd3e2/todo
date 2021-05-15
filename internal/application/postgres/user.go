package postgres

import (
	"github.com/jackc/pgx"
	"github.com/swd3e2/todo/internal/application"
)

// UserRepository Репозиторий пользователей
type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{conn}
}

func (r *UserRepository) FindByLogin(login string) (*application.User, error) {
	user := &application.User{}
	err := r.conn.
		QueryRow("select id, name, last_name, age, login, password, last_updated from users where login=$1", login).
		Scan(&user.Id, &user.Name, &user.LastName, &user.Age, &user.Login, &user.Password, &user.LastUpdated)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Save(user *application.User) error {
	isNewUser := user.Id == 0

	if isNewUser {
		if _, err := r.conn.Exec("insert into users(id, name, last_name, login, password) values(nextval('users_id_seq'), $1, $2, $3, $4)", user.Name, user.LastName, user.Login, user.Password); err != nil {
			return err
		}
	} else {
		if _, err := r.conn.Exec("update users set name=$1, last_name=$2, login=$3, password=$4 where login=$5", user.Name, user.LastName, user.Login, user.Password, user.Id); err != nil {
			return err
		}
	}

	return nil
}
