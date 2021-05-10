package application

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	userNotFound  = errors.New("пользователь не найден")
	wrongPassword = errors.New("неправильный пароль")
)

type Credentials struct {
	Login    string
	Password string
}

// User Пользователь
type User struct {
	Id          int64
	Name        string
	LastName    string
	Age         int
	LastUpdated time.Time
	Login       string
	Password    string
}

// UserRepository Репозиторий пользователей
type UserRepository interface {
	// FindByLogin Поиск пользователя по логину
	FindByLogin(login string) (*User, error)

	// Save Сохранение пользователя
	Save(user *User) error
}

// UserService Сервис пользователей
type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Token Jwt токен
type Token struct {
	Token string
}

// Authorize Авторизация пользователя, выдача токена
func (s *UserService) Authorize(login string, password string) (token *Token, err error) {
	var user *User
	fmt.Println(login)
	if user, err = s.userRepository.FindByLogin(login); err != nil || user == nil {
		return nil, userNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, wrongPassword
	}

	token = &Token{
		Token: "some_token",
	}

	return token, nil
}

// Register Регистрация нового пользователя
func (s *UserService) Register(Name string, LastName string, Login string, Password string) error {
	var bytes []byte
	var err error

	if bytes, err = bcrypt.GenerateFromPassword([]byte(Password), 14); err != nil {
		return err
	}

	password := string(bytes)

	user := &User{
		Name:     Name,
		LastName: LastName,
		Login:    Login,
		Password: password,
	}

	return s.userRepository.Save(user)
}
